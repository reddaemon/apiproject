package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/mmcdole/gofeed"
)

type Message struct {
	Title     string
	Link      string
	Published string
}

type Messages []Message

// GetNewsFeed method which get current rss state
func GetNewsFeed(w http.ResponseWriter, _ *http.Request) {
	fp := gofeed.NewParser()
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	feed, err := fp.ParseURLWithContext("https://www.sports.ru/stat/export/rss/taglenta.xml?id=1044512", ctx)
	if err != nil {
		log.Fatal("Cannot retrieve url and parse", err)
	}

	m := Message{}
	mm := Messages{}
	for _, i := range feed.Items {
		m.Title = i.Title
		m.Link = i.Link
		m.Published = i.Published
		mm = append(mm, m)
	}

	js, err := json.MarshalIndent(mm, " ", "   ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(js)
	if err != nil {
		log.Fatal("Cannot write response")
	}
}
