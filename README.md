# GoTutorial

## Setting up Go workspace
1. Create directory: mkdir GoTutorial
2. To inform Go tool about the workspace set the GOPATH environment variable: export GOPATH=~/Documents/GoWorkspace  (Can verify if the GOPATH is set or not through command "go env")
3. Add /usr/local/go/bin to the PATH environment variable: export PATH=$PATH:/usr/local/go/bin
4. Create src folder/hello and add hello.go inside that folder, add go code
5. To build using go tool, cd into src/hello and then : go build. The command above will build an executable named hello in the directory alongside your source code.
6. To run: ./hello
7. You can run go install to install the binary into your workspace's bin directory or go clean -i to remove it.

Reference: https://golang.org/doc/install 

Must Read: How to Write Go Code: https://golang.org/doc/code.html


## Useful Commands:
1. go build command compile packages and dependencies
2. go run command compiles and executes a program.
3. go install command compiles packages and creates binary executable files or package archive files.

    3.1 go install <package> command looks for any file with main package declaration inside given package directory. If it finds a file, then Go knows this is an executable program and it needs to create a binary file.
    
    3.2 If a package does not contain file with main package declaration, then Go creates a package archive (.a) file inside pkg directory. )
4. go doc prints documentation for a package
5. Other commands: https://golang.org/cmd/go/ 

## 
1. A Go workspace contains 3 major sub directories.
    1.1 src - contains source code files
    1.2 pkg - contains packages(libraries)
    1.3 bin - contains executables
2. Workspace directory is defined by the GOPATH variable