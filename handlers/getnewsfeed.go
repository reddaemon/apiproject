package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/mmcdole/gofeed"
)

// GetNewsFeed method which get current rss state
func GetNewsFeed(w http.ResponseWriter, _ *http.Request) {
	fp := gofeed.NewParser()
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	feed, err := fp.ParseURLWithContext("https://www.sports.ru/stat/export/rss/taglenta.xml?id=1044512", ctx)
	if err != nil {
		log.Fatal("Cannot retrieve url and parse", err)
	}
	for _, i := range feed.Items {
		fmt.Fprint(w, i.Title, i.Link, i.Published, "\n\n")
	}
}
