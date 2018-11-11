package main

import (
	"flag"
	"fmt"
)

// List dependencies of the given package
func dep() {
	packageName := flag.Arg(1)
	if packageName == "" {
		fmt.Print("Please specify a package name")
		return
	}

	packages, err := localIndex.ResolveDependencies(packageName, false)
	if err != nil {
		fmt.Printf("Failed to resolve dependencies: %s", err.Error())
		return
	}

	for _, depPackage := range packages {
		fmt.Println(depPackage)
	}
}

// List changes for the installed packages
func changes() {
	updates, err := localIndex.ResolveUpdates()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, pack := range updates {
		packOld := localIndex.InstalledPackages[pack]
		packNew := localIndex.RepositoryPackages[pack]
		fmt.Printf("v%d -> v%d\t%s\n", packOld.Version, packNew.Version, pack)
	}
}

// Search for a given package
func search() {
	keyword := flag.Arg(1)

	results := localIndex.Search(keyword)
	for _, pack := range results {
		fmt.Println(pack)
	}
}
