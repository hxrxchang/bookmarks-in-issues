package main

import (
	"log"

	"github.com/hxrxchang/bookmarks-in-issues/app"
)

func main() {
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
