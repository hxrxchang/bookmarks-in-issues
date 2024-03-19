package app

import (
	"fmt"
	"os"
	"strings"
)

func GetEnv() (string, *Repo, error) {
	token := os.Getenv("GITHUB_TOKEN")
	if token == "" {
		return "", nil, fmt.Errorf("GITHUB_TOKEN is required")
	}

	ghrepo := os.Getenv("GITHUB_REPOSITORY")
	if ghrepo == "" {
		return "", nil, fmt.Errorf("GITHUB_REPOSITORY is required")
	}

	splited := strings.Split(ghrepo, "/")
	if len(splited) != 2 {
		return "", nil, fmt.Errorf("GITHUB_REPOSITORY should be in the format of owner/repo")
	}
	repo := &Repo{
		Owner: splited[0],
		Name:  splited[1],
	}

	return token, repo, nil
}
