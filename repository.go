package pkg

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// Repository represents a repository
type Repository struct {
	Packages []RepositoryPackage
}

// FetchRepository fetches repository data from the given url
func FetchRepository(repositoryURL *url.URL) (*Repository, error) {
	response, err := http.Get(repositoryURL.String())
	if err != nil {
		return nil, fmt.Errorf("Error while trying to connect to repository %s: %s", repositoryURL.String(), err.Error())
	}

	data, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("Error while trying to read response: %s", err.Error())
	}

	repo := new(Repository)
	err = json.Unmarshal(data, &repo)
	if err != nil {
		return nil, fmt.Errorf("Error while parsing repository data: %s", err.Error())
	}

	return repo, nil
}
