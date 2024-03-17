package app

import (
	"flag"
	"fmt"
)

type Flags struct {
	URL   string
	Number int
}

func NewFlag() (*Flags, error) {
	url := flag.String("url", "", "URL written in the issue title")
	number := flag.Int("number", 0, "number of the issue")
	flag.Parse()

	if *url == "" {
		return nil, fmt.Errorf("URL is required")
	}

	if *number == 0 {
		return nil, fmt.Errorf("id is required")
	}

	return &Flags{
		URL: *url,
		Number: *number,
	}, nil
}
