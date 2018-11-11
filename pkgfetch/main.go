// A package management tool based on pkg
package main

import (
	"flag"
	"fmt"

	"github.com/benp98/pkg"
)

type actionHandler func()
type packageAction struct {
	description string
	handler     actionHandler
}

var localIndex *pkg.LocalIndex
var actions map[string]packageAction

func main() {
	flag.Parse()

	statusUpdater := pkg.LogStatusUpdater{}

	localIndex = pkg.NewLocalIndex("pkgfetch.json", statusUpdater)
	localIndex.Load()

	// Define a list of actions and their description
	actions = make(map[string]packageAction)
	actions["add"] = packageAction{"Add a repository to the list of active repositories", add}
	actions["update"] = packageAction{"Sync the local index with the active repositories", update}
	actions["install"] = packageAction{"Install a package", install}
	actions["upgrade"] = packageAction{"Upgrade installed packages", upgrade}
	actions["dep"] = packageAction{"List dependencies of a package", dep}
	actions["changes"] = packageAction{"List the packages which will be updated if running 'upgrade'", changes}
	actions["search"] = packageAction{"Search the local index for a given package", search}
	actions["help"] = packageAction{"Shows this help text", help}

	actionName := flag.Arg(0)
	item, exists := actions[actionName]
	if !exists {
		fmt.Println("Unknown action:", actionName)
		fmt.Println()
		help()
		return
	}

	item.handler()

	localIndex.Save()
}

// Prints help for the tool
func help() {
	fmt.Println("pkgfetch v0.1")
	fmt.Println("(c) 2018 Ben-Ove Pasinski. See https://github.com/benp98/pkg for more information.")
	fmt.Println()

	// Prints the list of actions and their description
	for k, v := range actions {
		fmt.Printf("%s\t%s\n", k, v.description)
	}
}
