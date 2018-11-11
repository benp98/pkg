package pkg

import (
	"archive/zip"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

// Update updates the local package index
func (index *LocalIndex) Update() error {
	index.RepositoryPackages = make(map[string]RepositoryPackage)

	for _, repoURL := range index.Repositories {
		url, err := url.Parse(repoURL)
		if err != nil {
			return err
		}

		repo, err := FetchRepository(url)
		if err != nil {
			return err
		}

		for _, pack := range repo.Packages {
			item, exists := index.RepositoryPackages[pack.Name]

			if (exists && item.Version < pack.Version) || !exists {
				index.RepositoryPackages[pack.Name] = pack
			}
		}
	}

	return nil
}

// Install looks at the local package index, fetches the associated archive and extracts it.
func (index *LocalIndex) Install(packageName string) error {
	index.statusUpdater.Message(fmt.Sprintf("Installing %s", packageName))
	item, exists := index.RepositoryPackages[packageName]
	if !exists {
		return fmt.Errorf("Package %s does not exist", packageName)
	}

	url, err := url.Parse(item.File)
	if err != nil {
		return fmt.Errorf("Invalid URL for package %s", packageName)
	}

	index.statusUpdater.Message(fmt.Sprintf("Starting download of %s", packageName))
	response, err := http.Get(url.String())
	if err != nil {
		return fmt.Errorf("Failed to download package %s from URL %s", packageName, url.String())
	}

	filename := "pkg_" + packageName + ".zip"
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("Failed to create temporary file %s", filename)
	}

	_, err = io.Copy(file, response.Body)
	if err != nil {
		return fmt.Errorf("Failed to download package %s", packageName)
	}
	file.Close()

	index.statusUpdater.Message(fmt.Sprintf("Unpacking %s", packageName))
	readCloser, err := zip.OpenReader(filename)
	if err != nil {
		return fmt.Errorf("Failed to open archive for %s", packageName)
	}

	for _, file := range readCloser.File {
		if file.FileInfo().IsDir() {
			os.MkdirAll(file.Name, 0777)
		} else {
			sourceFile, err := file.Open()
			if err != nil {
				return fmt.Errorf("Failed to read file %s from package %s", file.Name, packageName)
			}

			destFile, err := os.Create(file.Name)
			if err != nil {
				return fmt.Errorf("Failed to create file %s while installing %s", file.Name, packageName)
			}

			_, err = io.Copy(destFile, sourceFile)
			if err != nil {
				return fmt.Errorf("Failed to write file %s while installing %s", file.Name, packageName)
			}

			sourceFile.Close()
			destFile.Close()
		}
	}

	readCloser.Close()

	os.Remove(filename)

	index.InstalledPackages[packageName] = Package{item.Version}

	index.statusUpdater.Message(fmt.Sprintf("Package %s installed", packageName))

	return nil
}

// Upgrade looks for updates and installs them if available
func (index *LocalIndex) Upgrade() error {
	updates, err := index.ResolveUpdates()
	if err != nil {
		return fmt.Errorf("Could not resolve updates: %s", err.Error())
	}

	for _, pack := range updates {
		index.Install(pack)
	}

	return nil
}

// Search searches the local repository index for packages with the given name
func (index *LocalIndex) Search(keyword string) []string {
	list := make([]string, 0)

	for name := range index.RepositoryPackages {
		if strings.Contains(strings.ToLower(name), strings.ToLower(keyword)) {
			list = append(list, name)
		}
	}

	return list
}
