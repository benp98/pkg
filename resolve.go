package pkg

import (
	"fmt"
)

// ResolveDependencies creates a list of package dependencies
func (index *LocalIndex) ResolveDependencies(packageName string, ignoreInstalled bool) ([]string, error) {
	item, exists := index.RepositoryPackages[packageName]

	if !exists {
		return nil, fmt.Errorf("Package %s does not exist", packageName)
	}

	dependencies := make(map[string]bool)

	for _, dep := range item.Dependencies {
		_, isInstalled := index.InstalledPackages[dep]
		if ignoreInstalled && isInstalled {
			continue
		}

		dependencies[dep] = true

		innerDependencies, err := index.ResolveDependencies(dep, ignoreInstalled)
		if err != nil {
			return nil, fmt.Errorf("Failed to resolve dependencies: %s", err.Error())
		}

		for _, innerDependency := range innerDependencies {
			dependencies[innerDependency] = true
		}
	}

	list := make([]string, 0)

	for dependency := range dependencies {
		list = append(list, dependency)
	}

	return list, nil
}

// ResolveUpdates compares the versions of the installed packages with those from the repositories and creates a list of packages which have been updated
func (index *LocalIndex) ResolveUpdates() ([]string, error) {
	list := make([]string, 0)

	for pack, item := range index.InstalledPackages {
		repoItem, exists := index.RepositoryPackages[pack]

		if !exists {
			return nil, fmt.Errorf("Package %s does not exist (anymore?)", pack)
		}

		if repoItem.Version > item.Version {
			list = append(list, pack)
		}
	}

	return list, nil
}
