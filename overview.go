package main

/*
Each repository can contain multiple modules, and each module can contain multiple packages.

Multiple Modules in a Repository
   - Each module is a self-contained unit, typically represented by a go.mod file in its root directory.
   - A repository can have multiple directories, each with its own go.mod, thus creating multiple modules.
   - Modules are designed to manage dependencies, versioning, and deployment independently.

Multiple Packages in Each Module
   - Inside each module, you can have as many packages as you want. Each package is typically a directory with one or more .go files that share the same package name.
   - Packages are created by defining different directories (other than the module root) within the module.
   - In Go, each package can be imported by other packages in the same module or even in different modules, as long as they are appropriately referenced.


Example Directory Structure

my-repo/
│
├── module1/
│   ├── go.mod                    # Module 1 definition
│   ├── main.go                    # Main package
│   ├── packageA/                  # A separate package in module1
│   │   └── fileA.go
│   └── packageB/                  # Another package in module1
│       └── fileB.go
│
└── module2/
    ├── go.mod                    # Module 2 definition
    ├── main.go                    # Main package
    ├── packageC/                  # A separate package in module2
    │   └── fileC.go
    └── packageD/                  # Another package in module2
        └── fileD.go


 Usage Example

- Module 1 can import and use packageA and packageB within itself.
- Module 2 can import packageC and packageD within itself.
- If module2 needs to use packageA from module1, you can add it as a dependency in module2’s go.mod:

  go
  require example.com/my-repo/module1 v0.0.0

  replace example.com/my-repo/module1 => ../module1
*/