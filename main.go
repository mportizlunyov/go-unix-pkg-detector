// Written by Mikhail Ortiz-Lunyov
//
// Version 0.0.1

/*
Go UNIX package detector

This package works to get all of the package managers installed on the system,
and report them by by calling the report() method.

While not official intended to run as a stand-alone package,
more like a library to implement,
a main() method is provided to demonstrate the capacities of the program.
*/
package main

// Import packages
import (
	"flag"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

// Declare fields
// // Constants

// Constants related to versions
const (
	SHORT_VERSION_NUM string = "0.0.1"
	VERSION_NAME      string = "November 18th, 2024"
	DEV_MARKER        string = "-alpha"
	FULL_VERSION      string = "v" + SHORT_VERSION_NUM + DEV_MARKER + " ( " + VERSION_NAME + ")"
)

// Field containg the names of the official package managers to check for.
var officialPackageNameArray = [...]string{
	"apt",
	"yum",
	"transactional-update",
	"dnf",
	"rpm-ostree",
	"pacman",
	"apk",
	"zypper",
	"pacman",
	"xbps",
	"swupd",
	"slackpkg",
	"eopkg",
	"pkg",
	"pkg_add",
}

// Field containg the names of the alternative package managers to check for.
var alternativePackageNameArray = [...]string{
	"flatpak",
	"snap",
	"brew",
	"portsnap",
	"rubygem",
	"yarn",
	"pipx",
	"npm",
}

// Primary method used to return the detected package managers
func Report(pkgType int) ([]string, []string) {
	// TODO: Implement
	return nil, nil
}

// Helper method to delete duplicate values in a slice
func pruneSlice(sliceToPrune []string) []string {
	// Initialise variables
	var prunedSlice []string = sliceToPrune

	// Iterate through sliceToPrune, checking if
	// for index, item := range prunedSlice {
	// 	if index+1 < len(prunedSlice)-1 && prunedSlice[index+1] == item {
	// 		prunedSlice = append(prunedSlice[:index], prunedSlice[index+1:]...)
	// 	}
	// }

	slices.Sort(prunedSlice)
	prunedSlice = slices.Compact(prunedSlice)

	return prunedSlice
}

// Method used to check if directory contains binary with the name of a package manager listed above.
func checkPkgManBinaries(directory string, official bool) []string {
	// Initialise variables
	var returnPkgMan []string = []string{}

	// Read the contents of the directory, if possible
	entries, err := os.ReadDir(directory)
	if err != nil {
		log.Fatal(err)
	}

	// Iterate through contents of the directory, appending names appearing in the above fields
	for _, file := range entries {
		// Iterate through relevent lists, adding to returnPkgMan when needed
		switch official {
		case true:
			for _, pkgMan := range officialPackageNameArray {
				switch file.Name() {
				case pkgMan:
					returnPkgMan = append(returnPkgMan, file.Name())
				}
			}
		case false:
			for _, pkgMan := range alternativePackageNameArray {
				switch file.Name() {
				case pkgMan:
					returnPkgMan = append(returnPkgMan, file.Name())
				}
			}
		}
	}

	return returnPkgMan
}

// Method used to search through the user's path for names of package managers
func searchUserPATH() ([]string, []string) {
	// Initialise variables
	// // Get the path of the directories in $PATH variable
	var pathDirectories []string = strings.Split(os.Getenv("PATH"), ":")
	// // Get the names of identified package managers from a PATH directory
	var identifiedOfficialPkgMan []string = []string{}
	var identifiedAlternativePkgMan []string = []string{}

	// Iterate through user's $PATH variable, checking the binaries that exist
	for _, directory := range pathDirectories {
		identifiedOfficialPkgMan = append(identifiedOfficialPkgMan, checkPkgManBinaries(directory, true)...)
		identifiedAlternativePkgMan = append(identifiedAlternativePkgMan, checkPkgManBinaries(directory, false)...)
	}

	return pruneSlice(identifiedOfficialPkgMan), pruneSlice(identifiedAlternativePkgMan)
}

// Main method to demonstrate capabilities of this package.
func main() {
	// Check flags
	// // -v / --version
	versionShort := flag.Bool("v", false, "Print the version number of this module")
	versionLong := flag.Bool("version", false, "Long form of [-v]")
	// // Parse flags
	flag.Parse()
	// // // Finalise flags
	versionFlag := *versionShort || *versionLong
	// Print version and exit if requested
	if versionFlag {
		fmt.Println(FULL_VERSION)
		os.Exit(0)
	}

	fmt.Println(searchUserPATH())
}
