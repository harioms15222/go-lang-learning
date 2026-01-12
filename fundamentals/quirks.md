## Note: undefined function across files

- Cause: running `go run main.go` compiles only the listed file(s); functions defined in other files (e.g., `functions.go`) are excluded, so `Functions` is undefined.
- Fix: run the package (recommended) with `go run .`, or `go run *.go`, or explicitly include all files: `go run main.go functions.go`.
- Tip: files in the same package (e.g., `package main`) can call each other as long as you compile/run all files together.

## modules 
go mod init go.mod starts a file and then we can run `go run .` just `main.go` should be there