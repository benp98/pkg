// Package pkg provides types and methods for a simple package management system
package pkg

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// LocalIndex stores the state of the local package index
type LocalIndex struct {
	fileName           string
	statusUpdater      StatusUpdater
	Repositories       []string
	InstalledPackages  map[string]Package
	RepositoryPackages map[string]RepositoryPackage
}

// NewLocalIndex initializes a new local index instance
func NewLocalIndex(fileName string, statusUpdater StatusUpdater) *LocalIndex {
	index := new(LocalIndex)
	index.fileName = fileName
	index.statusUpdater = statusUpdater

	index.Repositories = make([]string, 0)
	index.InstalledPackages = make(map[string]Package)
	index.RepositoryPackages = make(map[string]RepositoryPackage)

	return index
}

// Load loads the local index from disk
func (index *LocalIndex) Load() error {
	file, err := os.Open(index.fileName)
	if err != nil {
		return fmt.Errorf("Failed to load local package index: %s", err.Error())
	}

	data, err := ioutil.ReadAll(file)
	file.Close()
	if err != nil {
		return fmt.Errorf("Failed to load local package index: %s", err.Error())
	}

	err = json.Unmarshal(data, index)

	return nil
}

// Save stores the local package index to disk
func (index *LocalIndex) Save() error {
	file, err := os.Create(index.fileName)
	if err != nil {
		return fmt.Errorf("Failed to save local package index: %s", err.Error())
	}

	data, err := json.Marshal(index)
	if err != nil {
		return fmt.Errorf("Failed to save local package index: %s", err.Error())
	}

	file.Write(data)
	file.Close()

	return nil
}
