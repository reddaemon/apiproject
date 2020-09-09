package handlers

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/mmcdole/gofeed"

	"context"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/reddaemon/apiproject/config"
	database "github.com/reddaemon/apiproject/db"
	"github.com/reddaemon/apiproject/internal/database/postgres"
	"github.com/reddaemon/apiproject/logger"
)

type PsqlRepository struct {
	*sqlx.DB //nolint
}

var ConfigPath = flag.String("config", "config.yml", "path to config file") //nolint

// PuttoDb is handler which call insert to DB method

func PutToDb(w http.ResponseWriter, _ *http.Request) {
	c, err := config.GetConfig(*ConfigPath)
	if err != nil {
		log.Fatalf("unable to get config: %v", err)
	}
	l, err := logger.GetLogger(c)
	if err != nil {
		log.Fatalf("unable to get logger: %v", err)
	}
	db, err := database.GetDb(c)
	if err != nil {
		log.Fatalf("unable to get db: %v", err)
	}

	r := postgres.NewPsqlRepository(db, l)
	ctx := context.Background()
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL("https://www.sports.ru/stat/export/rss/taglenta.xml?id=1044512")
	if err != nil {
		log.Fatal("Cannot retrieve url and parse", err)
	}
	for _, i := range feed.Items {
		err = postgres.PsqlRepository.InsertToDB(r, ctx, feed, i.Title, i.Link, i.Published)
		fmt.Fprint(w, "Done")
		if err != nil {
			return
		}
	}
	fmt.Fprint(w, "Done")

}
