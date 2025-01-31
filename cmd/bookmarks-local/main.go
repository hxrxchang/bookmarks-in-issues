package main

import (
	"flag"
	"log"
	"os"
	"os/exec"

	"github.com/hxrxchang/bookmarks-in-issues/app"
)

// GitHub Actionsからアクセスできないページ用の、ローカルで動かすCLI
// gh (GitHub CLI) をインストールして、gh auth login で認証しておく
func main() {
	url := flag.String("url", "", "URL")
	flag.Parse()

	if *url == "" {
		log.Fatal("url is required")
	}

	title, err := app.FetchTitle(*url)
	if err != nil {
		log.Fatal(err)
	}

	if err := createGitHubIssue(title, *url); err != nil {
		log.Fatal(err)
	}

	log.Println("Issue created successfully")
}

func createGitHubIssue(title, description string) error {
	cmd := exec.Command("gh", "issue", "create", "--title", title, "--body", description)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
