# Next Language 🚀

Next is a powerful Generic **Interface Definition Language (IDL)** designed to create highly customized code across multiple programming languages. It leverages a flexible template system to transform high-level specifications into language-specific implementations.

[Language Specification (English)](./docs/en/language_spec.md) | [语言规范 (中文)](./docs/zh/language_spec.md)

## ✨ Key Features

- 🌐 Multi-language code generation from a single source
- 📝 Powerful templating system based on Go's text/template
- 🧩 Flexible customization through template inheritance and overloading
- 🏗️ Rich type system supporting interfaces, structs, enums, and various data types
- 🏷️ Annotation support for metadata and customization

## 🛠️ Template System

Next uses a template system based on Go's [text/template](https://pkg.go.dev/text/template/) package, with custom enhancements for code generation. Templates in Next use the `.npl` file extension.

### 📊 Template Hierarchy and Inheritance

Next implements a three-layer template hierarchy, allowing for easy customization and overriding of default behaviors:

1. Next builtin base templates: `next/<object_type>`
2. Next builtin language-specific templates: `next/<lang>/<object_type>`
3. User-defined language-specific templates: `<lang>/<object_type>`

This hierarchy enables a powerful inheritance and overloading mechanism, similar to class inheritance in object-oriented programming.

When rendering a template, Next searches for the most specific template first (user-defined), then falls back to language-specific templates, and finally to base templates if no overrides are found.

## 📥 Installation

### Unix-like Systems (Linux, macOS, Git Bash (for Windows) etc.)

To install Next on Unix-like systems, you can use the following command:

```sh
curl -fsSL https://getnext.sh | sh
```

This script will download and install the latest version of Next on your system.

### Windows

For Windows users, follow these steps:

1. Download the installation package from [dl.getnext.sh](https://dl.getnext.sh).
2. Extract the downloaded ZIP file.
3. Right-click on the `install.bat` file in the extracted folder.
4. Select "Run as administrator" to execute the installation script.

This will install the latest version of Next on your Windows system.

### Installing from Source

To install Next from source, use the following command:

```sh
go install github.com/next/next@latest
```

Make sure you have [Go](https://go.dev) `1.21+` installed on your system before running this command.

## 📚 Built-in Language Support

Next includes built-in templates for various languages, including C++, Go, Java, and more. These serve as a foundation for code generation and can be easily customized or extended using the template hierarchy system.

## 📖 Documentation

For detailed information on Next's syntax, features, and usage, please visit the official website at [nextlang.org](https://nextlang.org). The website provides comprehensive documentation, tutorials, and examples to help you get started with Next and make the most of its capabilities.

## 📝 TODO

- [ ] Implement annotation solver functionality to call external programs for calculating annotation parameters
- [ ] Implement vscode extensions to support Next source and Next template
- [ ] Handle template call stack for traceable error output
- [ ] Implement consistency check for enum member types and provide type of enum value
- [ ] Implement import dirs for importing files
- [ ] Create comprehensive user documentation, supporting both Chinese and English
- [ ] Support more built-in templates for additional programming languages

## 🤝 Contributing

We welcome contributions to Next Language! Please see our [Contribution Guidelines](CONTRIBUTING.md) for more information on how to get started.

## 📄 License

Next is released under the [MIT License](LICENSE).
