package main

import (
	"flag"
	"fmt"
)

func install() {
	packageName := flag.Arg(1)
	if packageName == "" {
		fmt.Print("Please specify a package to install")
		return
	}

	packages, err := localIndex.ResolveDependencies(packageName, true)
	if err != nil {
		fmt.Printf("Failed to resolve dependencies: %s", err.Error())
		return
	}

	packages = append(packages, packageName)

	for _, installPackage := range packages {
		//fmt.Println("Installing", installPackage)
		err = localIndex.Install(installPackage)
		if err != nil {
			fmt.Printf("Failed to install package: %s", err.Error())
			return
		}
	}
}

func upgrade() {
	err := localIndex.Upgrade()
	if err != nil {
		fmt.Println(err.Error())
	}
}
