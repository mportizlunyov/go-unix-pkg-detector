# UNIX Package Manager Detector (Go edition)
This Go library is a simple tool to detect and return all package managers installed on a system.

To use it in your project, run the following command: `go get github.com/mportizlunyov/go-unix-pkg-detector`

Then, import `github.com/mportizlunyov/go-unix-pkg-detector` in your Go program.
See the following methods below.

### `Report() ([]string, []string)`
Returns 2 []string arrays containing:
 1. Official, distribution package managers
 2. Alternative package managers

### `Version() (string)`
Returns the full version, including version name, of this library

### Example:
Code:
```
package main

import (
	"fmt"

	// Imported library
	"github.com/mportizlunyov/go-unix-pkg-detector/unixpkgdetector"
)

func main() {
	fmt.Println("Hello!")
	// Report() method
	fmt.Println(unixpkgdetector.Report())
	// Version() method
	fmt.Println(unixpkgdetector.Version())
}
```
Result on a typical default Ubuntu systems:
```
Hello!
[apt] [snap]
v1.0.4-release ( November 24th, 2024)
```
