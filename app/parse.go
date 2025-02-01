package app

import (
	"net/url"

	"github.com/gocolly/colly/v2"
)


type InvalidUrlError struct {
	URL string
}
func (e *InvalidUrlError) Error() string {
	return e.URL + " is invalid URL"
}

func IsValidUrl(URL string) error {
	parsedURL, err := url.Parse(URL)
	if err != nil || parsedURL.Scheme == "" || parsedURL.Host == "" {
		return &InvalidUrlError{URL: URL}
	}
	return nil
}

func FetchTitle(URL string) (string, error) {
	c := colly.NewCollector()
	var title string
	c.OnHTML("head title", func(e *colly.HTMLElement) {
		title = e.Text
	})
	err := c.Visit(URL)

	if err != nil {
		return "", err
	}

	return title, nil
}
