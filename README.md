# Learning Go

This repository contains my journey in learning the Go programming language. It includes example programs and explanations of Go concepts.

## Installation

To get started with Go:

1. Download and install Go from the [official Go website](https://golang.org/dl/).
2. Verify your installation by opening a terminal and running:
   ```
   go version
   ```
   This should display the installed Go version.

3. Set up your Go workspace (optional for Go 1.11+):
   - Set the `GOPATH` environment variable to your desired workspace location.
   - Add `$GOPATH/bin` to your `PATH`.

## Running Go Programs

### Example: /packages/main.go

To run the main program in the `packages` directory:

1. Navigate to the repository root.
2. Run the following command:
   ```
   go run ./packages/main.go
   ```

This will compile and run the program without producing an executable file.

Alternatively, you can build and then run the executable:

```
go build ./packages/main.go
./main  # On Windows, use .\main.exe
```

## Understanding go.work

The `go.work` file is used in Go 1.18+ for managing multi-module workspaces. It allows you to work on multiple related modules simultaneously.

### Structure of go.work

A typical `go.work` file might look like this:

```
go 1.18

use (
    ./packages
    ./anothermodule
)
```

This tells Go to use the modules in the `packages` and `anothermodule` directories.

### Using go.work

1. Create a `go.work` file in your root directory:
   ```
   go work init ./packages ./anothermodule
   ```

2. To add a new module later:
   ```
   go work use ./newmodule
   ```

3. Run programs using module paths:
   ```
   go run example.com/mymodule/packages
   ```

## Running Go Programs from CLI

1. Run a single file:
   ```
   go run filename.go
   ```

2. Run a package:
   ```
   go run .
   ```
   (When in the package directory)

3. Run with arguments:
   ```
   go run main.go arg1 arg2
   ```

4. Build and run:
   ```
   go build
   ./executable_name
   ```

5. Install and run (adds the executable to `$GOPATH/bin`):
   ```
   go install
   executable_name
   ```

Remember to run `go mod tidy` in each module directory to manage dependencies.

## Next Steps

Explore the various packages and programs in this repository. Each directory contains examples and exercises to help you learn Go concepts.

Happy coding!