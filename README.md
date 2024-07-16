# Purpose
Simple Go program for creating nested dummy directories and files for testing. Current functionality is pretty basic and requires editing `main.go` to change the count of nested directories (`directoryCount`).

Separate directories are created in a `topDirectory`. Directory names are numbers indicating how many directories are nested within it. A single `file.txt` is made in each directory with `fileContent` contents.

The top directory can be moved to whatever location needed for testing purposes.

# Requirements
- [Go](https://go.dev/)

# Instructions
- Clone this repo.
- Change `directoryCount` as needed within `main.go`.
- Run `go run main.go` within the top level of this project. This will create the `tmp` directory containing the nested directories.

# Future improvements
- Allow `directoryCount` to be an argument?
- Specify a lower limit to how many nested directories are created (ex. allow making a single nested directory with a depth of 20).