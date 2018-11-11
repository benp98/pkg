package main

import (
	"flag"
	"fmt"
	"net/url"
)

func add() {
	repoURL := flag.Arg(1)
	if repoURL == "" {
		fmt.Println("Please specify a repository URL")
		return
	}

	_, err := url.Parse(repoURL)
	if err != nil {
		fmt.Println("Please specify a valid repository URL")
		return
	}

	localIndex.Repositories = append(localIndex.Repositories, repoURL)
}

func update() {
	err := localIndex.Update()
	if err != nil {
		fmt.Printf("Failed to update the local index: %s", err.Error())
		return
	}
}
