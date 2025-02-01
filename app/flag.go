package app

import (
	"flag"
	"fmt"
	"log"
)

type Flags struct {
	URL    string
	Number int
}

func NewFlag() (*Flags, error) {
	url := flag.String("url", "", "URL written in the issue title")
	number := flag.Int("number", 0, "number of the issue")
	flag.Parse()

	log.Printf("url: %s, number: %d\n", *url, *number)

	if *url == "" {
		return nil, fmt.Errorf("url is required")
	}

	if *number == 0 {
		return nil, fmt.Errorf("nuber is required")
	}

	return &Flags{
		URL:    *url,
		Number: *number,
	}, nil
}
