package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	url := flag.String("url", "", "URL written in the issue title")
	flag.Parse()
	title, err := fetchTitle(*url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetchTitle: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(title)
}

func fetchTitle(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return "", err
	}

	var f func(*html.Node) string
	f = func(n *html.Node) string {
		if n.Type == html.ElementNode && n.Data == "title" {
			return n.FirstChild.Data
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			result := f(c)
			if result != "" {
				return result
			}
		}
		return ""
	}

	title := f(doc)
	if title == "" {
		return "", fmt.Errorf("no title element found")
	}

	return title, nil
}
