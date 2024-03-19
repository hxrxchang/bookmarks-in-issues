package app

import "github.com/gocolly/colly/v2"

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
