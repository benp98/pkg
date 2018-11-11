package pkg

// Package is a generic package type
type Package struct {
	Version uint64
}

// RepositoryPackage is the data structure which represents a package from a repository
type RepositoryPackage struct {
	Package

	Name         string
	File         string
	Dependencies []string
}
