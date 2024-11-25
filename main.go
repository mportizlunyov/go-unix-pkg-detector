// Written by Mikhail Ortiz-Lunyov
//
// Version 1.0.3-release

/*
Go UNIX package detector

This package works to get all of the package managers installed on the system,
and report them by by calling the report() method.

While not official intended to run as a stand-alone package,
more like a library to implement,
a main() method is provided to demonstrate the capacities of the program.

When getting the package managers, one needs to call the Report() method to get
both official and alternative package managers.
*/
package main

// Import packages
import (
	"log"
	"os"
	"strings"
)

// Declare fields
// // Constants

// Constants related to versions
const (
	short_VERSION_NUM string = "1.0.3"
	version_NAME      string = "November 24th, 2024"
	dev_MARKER        string = "-release"
	full_VERSION      string = "v" + short_VERSION_NUM + dev_MARKER + " ( " + version_NAME + ")"
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
func Report() ([]string, []string) {
	// Initialises variables
	var officialPkgsCollected []string
	var alternativePkgsCollected []string

	// Get results from $PATH, and append to internal slices
	collectedOffTemp, collectedAltTemp := searchUserPATH()
	officialPkgsCollected = append(officialPkgsCollected, collectedOffTemp...)
	alternativePkgsCollected = append(collectedAltTemp, collectedAltTemp...)

	// Return collected package managers, pruning duplicates
	return pruneSlice(officialPkgsCollected), pruneSlice(alternativePkgsCollected)
}

// Helper method to check if an item exists in a slice
func contains(sliceToInvestigage []string, searchItem string) bool {
	for _, item := range sliceToInvestigage {
		if searchItem == item {
			return true
		}
	}

	return false
}

// Helper method to delete duplicate values in a slice
func pruneSlice(sliceToPrune []string) []string {
	// Initialise variables
	var prunedSliceToReturn []string

	// Iterate through sliceToPrune, adding unique variables
	for _, item := range sliceToPrune {
		if len(prunedSliceToReturn) == 0 || !contains(sliceToPrune, item) {
			prunedSliceToReturn = append(prunedSliceToReturn, item)
		} else {
			continue
		}
	}

	return prunedSliceToReturn
}

// Method used to check if directory contains binary with the name of a listed package manager
func checkPkgManBinaryName(directory string, official bool) []string {
	// Initialise variables
	var returnPkgMan []string
	var referencePkgManList []string

	// Read the contents of the directory, if possible
	entries, err := os.ReadDir(directory)
	if err != nil {
		log.Fatal(err)
	}

	// Set the reference list
	switch official {
	case true:
		referencePkgManList = officialPackageNameArray[:]
	case false:
		referencePkgManList = alternativePackageNameArray[:]
	}

	// Iterate through contents of the directory, appending names appearing in the above fields
	for _, file := range entries {
		// // Iterate through relevent lists, adding to returnPkgMan when needed
		for _, pkgMan := range referencePkgManList {
			switch file.Name() {
			case pkgMan:
				returnPkgMan = append(returnPkgMan, file.Name())
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
	var identifiedOfficialPkgMan []string
	var identifiedAlternativePkgMan []string

	// Iterate through user's $PATH variable, checking the binaries that exist
	for _, directory := range pathDirectories {
		identifiedOfficialPkgMan = append(identifiedOfficialPkgMan, checkPkgManBinaryName(directory, true)...)
		identifiedAlternativePkgMan = append(identifiedAlternativePkgMan, checkPkgManBinaryName(directory, false)...)
	}

	return identifiedOfficialPkgMan, identifiedAlternativePkgMan
}
