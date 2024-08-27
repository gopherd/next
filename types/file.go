package types

import (
	"github.com/gopherd/next/ast"
	"github.com/gopherd/next/token"
)

// File represents a Next source file.
type File struct {
	pos        token.Pos
	pkg        *Package
	unresolved struct {
		annotations *ast.AnnotationGroup
	}

	Path        string
	Doc         *CommentGroup
	Annotations *AnnotationGroup

	decls   []*Decl
	stmts   []Stmt
	imports *Imports

	// all symbols used in current file:
	// - values: constant, enum member
	// - types: struct, enum
	symbols map[string]Symbol
}

func newFile(ctx *Context, src *ast.File) *File {
	file := &File{
		pos:     src.Pos(),
		Doc:     newCommentGroup(src.Doc),
		imports: &Imports{},
		symbols: make(map[string]Symbol),
	}
	file.unresolved.annotations = src.Annotations
	for _, d := range src.Decls {
		file.decls = append(file.decls, newDecl(ctx, file, d.(*ast.GenDecl)))
	}
	for _, s := range src.Stmts {
		file.stmts = append(file.stmts, newStmt(ctx, file, s))
	}
	if pos, err := file.createSymbols(); err != nil {
		ctx.errors.Add(ctx.fset.Position(pos), err.Error())
	}
	return file
}

func (f *File) Decls() []*Decl {
	var hasImport bool
	for _, d := range f.decls {
		if d.tok == token.IMPORT {
			hasImport = true
			break
		}
	}
	if !hasImport {
		return f.decls
	}
	var decls = make([]*Decl, 0, len(f.decls))
	for _, d := range f.decls {
		if d.tok != token.IMPORT {
			decls = append(decls, d)
		}
	}
	return decls
}

func (f *File) Stmts() []Stmt { return f.stmts }

func (f *File) Imports() *Imports { return f.imports }

func (f *File) Consts() []*ValueSpec {
	var consts []*ValueSpec
	for _, d := range f.decls {
		if d.tok == token.CONST {
			for _, s := range d.Specs {
				consts = append(consts, s.(*ValueSpec))
			}
		}
	}
	return consts
}

func (f *File) Enums() []*EnumType {
	var enums []*EnumType
	for _, d := range f.decls {
		if d.tok == token.ENUM {
			for _, s := range d.Specs {
				enums = append(enums, s.(*EnumType))
			}
		}
	}
	return enums
}

func (f *File) Structs() []*StructType {
	var structs []*StructType
	for _, d := range f.decls {
		if d.tok == token.STRUCT {
			for _, s := range d.Specs {
				structs = append(structs, s.(*StructType))
			}
		}
	}
	return structs
}

// LookupLocalType looks up a type by name in the file's symbol table.
// If the type is not found, it returns an error. If the symbol
// is found but it is not a type, it returns an error.
func (f *File) LookupLocalType(name string) (Type, error) {
	return expectTypeSymbol(name, f.symbols[name])
}

// LookupLocalValue looks up a value by name in the file's symbol table.
// If the value is not found, it returns an error. If the symbol
// is found but it is not a value, it returns an error.
func (f *File) LookupLocalValue(name string) (*ValueSpec, error) {
	return expectValueSymbol(name, f.symbols[name])
}

func (f *File) addSymbol(name string, s Symbol) error {
	if prev, dup := f.symbols[name]; dup {
		return &SymbolRedefinedError{Name: name, Prev: prev}
	}
	f.symbols[name] = s
	return nil
}

func (f *File) createSymbols() (token.Pos, error) {
	for _, d := range f.decls {
		for _, s := range d.Specs {
			switch s := s.(type) {
			case *ImportSpec:
				f.imports.Specs = append(f.imports.Specs, s)
			case *ValueSpec:
				if err := f.addSymbol(s.name, s); err != nil {
					return s.pos, err
				}
			case *EnumType:
				if err := f.addSymbol(s.name, s); err != nil {
					return s.pos, err
				}
				for _, m := range s.Members {
					if err := f.addSymbol(joinSymbolName(s.name, m.name), m); err != nil {
						return m.pos, err
					}
				}
			case *StructType:
				if err := f.addSymbol(s.name, s); err != nil {
					return s.pos, err
				}
			}
		}
	}
	return token.NoPos, nil
}

func (f *File) resolve(ctx *Context) {
	f.Annotations = ctx.resolveAnnotationGroup(f, f.unresolved.annotations)
	for _, d := range f.decls {
		d.resolve(ctx, f, f)
	}
}
