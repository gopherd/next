package main

import (
	"bytes"
	"embed"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"slices"
	"strconv"
	"strings"
	"text/template"

	"github.com/gopherd/core/builder"
	"github.com/gopherd/core/flags"
	"github.com/gopherd/core/term"
	"github.com/next/next/src/fsutil"
	"github.com/next/next/src/parser"
	"github.com/next/next/src/scanner"
	"github.com/next/next/src/types"
)

//go:embed builtin/*
var builtin embed.FS

const currentDir = "."
const nextExt = ".next"
const website = "https://next.as"
const repository = "https://github.com/next/next"

func exec(platform types.Platform, args []string) {
	stdin, stderr := platform.Stdin(), platform.Stderr()
	if len(args) == 2 && args[1] == "version" {
		builder.PrintInfo()
		unwrap(stderr, 0)
	}

	flagSet := flag.NewFlagSet(args[0], flag.ContinueOnError)
	flagSet.Usage = func() {}

	compiler := types.NewCompiler(platform, builtin)
	compiler.SetupCommandFlags(flagSet, flags.UseUsage(flagSet.Output(), flags.NameColor(term.Bold)))

	// set output color for error messages
	flagSet.SetOutput(term.ColorizeWriter(stderr, term.Red))
	usageFunc := func() {
		flagSet.SetOutput(stderr)
		name := term.Bold.Colorize(args[0])
		term.Fprint(flagSet.Output(), "Next is an IDL for generating customized code across multiple languages.\n\n")
		term.Fprint(flagSet.Output(), "Usage:\n")
		term.Fprintf(flagSet.Output(), "  %s [Options] [source_dirs_or_files...] (default: current directory)\n", name)
		term.Fprintf(flagSet.Output(), "  %s [Options] <stdin>\n", name)
		term.Fprintf(flagSet.Output(), "  %s version\n", name)
		term.Fprintf(flagSet.Output(), "\nOptions:\n")
		flagSet.PrintDefaults()
		term.Fprintf(flagSet.Output(), `For more information:
  Website:    %s
  Repository: %s

`,
			(term.Bold + term.BrightBlue).Colorize(website),
			(term.Bold + term.BrightBlue).Colorize(repository),
		)
	}
	if err := flagSet.Parse(args[1:]); err != nil {
		if err == flag.ErrHelp {
			usageFunc()
			unwrap(stderr, 0)
		}
		usage(flagSet, stderr, "")
	}

	var files []string
	var source io.Reader
	if flag.NArg() == 0 {
		if osStdin, ok := stdin.(*os.File); ok {
			if stat, err := osStdin.Stat(); err == nil && (stat.Mode()&os.ModeCharDevice) == 0 {
				source = osStdin
				files = append(files, "<stdin>")
			} else {
				v, e := fsutil.AppendFiles(files, currentDir, nextExt, false)
				files = result(stderr, v, e)
			}
		} else if stdin != nil {
			source = stdin
			files = append(files, "<input>")
		}
	} else {
		for _, arg := range flag.Args() {
			if arg == "-" {
				usage(flagSet, stderr, "invalid argument: -")
			}
			if strings.HasPrefix(arg, "-") {
				usage(flagSet, stderr, "flag %q not allowed after non-flag argument", arg)
			}
		}
		for _, arg := range flag.Args() {
			v, e := fsutil.AppendFiles(files, arg, nextExt, false)
			files = result(stderr, v, e)
		}
	}
	if len(files) == 0 {
		usage(flagSet, stderr, "no source files")
	}

	// compute absolute path and remove duplicated files
	if source == nil {
		seen := make(map[string]bool)
		for i := len(files) - 1; i >= 0; i-- {
			v, e := filepath.Abs(files[i])
			files[i] = result(stderr, v, e)
			if seen[files[i]] {
				files = append(files[:i], files[i+1:]...)
			} else {
				seen[files[i]] = true
			}
		}
	}

	// parse and resolve all files
	for _, file := range files {
		if source == nil {
			v, e := platform.ReadFile(file)
			source = bytes.NewReader(result(stderr, v, e))
		}
		v, e := parser.ParseFile(compiler.FileSet(), file, source, parser.ParseComments)
		f := result(stderr, v, e)
		v2, e2 := compiler.AddFile(f)
		result(stderr, v2, e2)
	}
	unwrap(stderr, compiler.Resolve())

	// generate files
	unwrap(stderr, types.Generate(compiler))
}

// exit exits the program. It is a variable for overriding.
var exit = os.Exit

// usage prints usage message and exits the program. It is a variable for overriding.
var usage = func(flagSet *flag.FlagSet, stderr io.Writer, format string, args ...any) {
	if format != "" {
		fmt.Fprintln(flagSet.Output(), fmt.Sprintf(format, args...))
	}
	unwrap(stderr, fmt.Errorf("try %q for help", flagSet.Name()+" -h"))
}

// result returns the value if err is nil, otherwise it unwraps the error and exits the program.
func result[T any](stderr io.Writer, v T, err error) T {
	unwrap(stderr, err)
	return v
}

// unwrap unwraps error and exits the program if error is not nil.
func unwrap(stderr io.Writer, err any) {
	switch err := err.(type) {
	case nil:
		return // do nothing if no error
	case error:
		if errs, ok := err.(scanner.ErrorList); ok {
			const maxErrorCount = 20
			for i := 0; i < len(errs) && i < maxErrorCount; i++ {
				printErrorWithPosition(stderr, errs[i].Error())
			}
			if remaining := len(errs) - maxErrorCount; remaining > 0 {
				fmt.Fprintf(stderr, "and %d more errors\n", remaining)
			}
		} else {
			printError(stderr, err)
		}
		exit(2)
	case string:
		fmt.Fprintln(stderr, err)
		exit(1)
	case int:
		exit(err)
	default:
		fmt.Fprintln(stderr, err)
		exit(1)
	}
}

func printError(stderr io.Writer, err error) {
	if err == nil {
		return
	}
	var errs []string
	for err != nil {
		if e, ok := err.(template.ExecError); ok {
			errs = append(errs, err.Error())
			err = e.Err
		} else if e := errors.Unwrap(err); e != nil {
			err = e
		} else {
			errs = append(errs, err.Error())
			break
		}
	}
	for i := 0; i+1 < len(errs); i++ {
		errs[i] = strings.TrimSuffix(strings.TrimSpace(errs[i]), strings.TrimSpace(errs[i+1]))
	}
	slices.Reverse(errs)
	for i := len(errs) - 1; i >= 0; i-- {
		s := strings.TrimSpace(errs[i])
		s = strings.TrimPrefix(s, "template: ")
		s = strings.TrimSuffix(s, ":")
		if s == "" || strings.HasPrefix(parseFilename(s), types.StubPrefix) {
			errs = append(errs[:i], errs[i+1:]...)
		} else {
			errs[i] = s
		}
	}
	if len(errs) == 0 {
		printErrorWithPosition(stderr, err.Error())
		return
	}
	printErrorWithPosition(stderr, errs[0])
	maxIndent := 32
	if len(errs) > maxIndent {
		maxIndent = 0
	}
	for i := 1; i < len(errs); i++ {
		fmt.Fprint(stderr, strings.Repeat(" ", min(i, maxIndent)))
		printErrorWithPosition(stderr, errs[i])
	}
}

func parseFilename(err string) string {
	parts := strings.SplitN(err, ":", 2)
	if len(parts) < 2 {
		return ""
	}
	return parts[0]
}

// printErrorWithPosition tries to print template error in a more readable format.
// template error format: "<filename>:<line>:<column>: <error message>"
func printErrorWithPosition(stderr io.Writer, err string) {
	const fileColor = term.Color("")
	const lineColor = term.BrightBlue
	const columnColor = term.BrightGreen
	const errorColor = term.BrightRed

	if err == "" {
		return
	}
	parts := strings.SplitN(err, ":", 4)
	if len(parts) < 3 {
		term.Fprintln(stderr, errorColor.Colorize(err))
		return
	}
	filename := parts[0]
	if wd, err := os.Getwd(); err == nil {
		if rel, err := filepath.Rel(wd, filename); err == nil {
			filename = rel
		}
	}
	line := parts[1]
	column := ""
	if len(parts) > 3 {
		column = parts[2]
	}
	if !strings.HasSuffix(filename, nextExt) {
		// add 1 to column if it is a number for non-next files (most likely for template files)
		if i, err := strconv.Atoi(column); err == nil {
			column = strconv.Itoa(i + 1)
		}
	}
	message := parts[len(parts)-1]
	if column == "" {
		term.Fprintf(
			stderr, "%s:%s:%s\n",
			fileColor.Colorize(filename),
			lineColor.Colorize(line),
			errorColor.Colorize(message),
		)
	} else {
		term.Fprintf(
			stderr, "%s:%s:%s:%s\n",
			fileColor.Colorize(filename),
			lineColor.Colorize(line),
			columnColor.Colorize(column),
			errorColor.Colorize(message),
		)
	}
}
