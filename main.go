package main

import (
	"errors"
	"log"

	"github.com/hxrxchang/bookmarks-in-issues/app"
)

func main() {
	if err := app.Run(); err != nil {
		var invalidUrlError *app.InvalidUrlError
		if errors.As(err, &invalidUrlError) {
			// 正常終了させる
			log.Println(err)
		} else {
			log.Fatalln(err)
		}
	}
}
