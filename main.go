package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/reddaemon/apiproject/config"
	lh "github.com/reddaemon/apiproject/handlers"
)

func main() {
	flag.Parse()
	c, _ := config.GetConfig(*lh.ConfigPath)

	r := mux.NewRouter()

	srv := &http.Server{
		Addr:         c.URL,
		WriteTimeout: time.Second * c.Writetimeout,
		ReadTimeout:  time.Second * c.Readtimeout,
		IdleTimeout:  time.Second * c.Idletimeout,
		Handler:      r,
	}

	r.Handle("/", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(lh.HomeHandler)))
	r.Handle("/news", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(lh.GetNewsFeed)))
	r.Handle("/puttodb", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(lh.PutToDb)))

	log.Printf("Server started and serving on port %s", c.URL)

	// Run server in goroutine
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	ch := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT(CTRL+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(ch, os.Interrupt)

	//Block until we receive out signal
	<-ch

	// Create a deadline to wait for
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*c.Gracefultimeout)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline

	go func() {
		err := srv.Shutdown(ctx)
		if err != nil {
			log.Fatal("server is crash")
		}
		<-ctx.Done()
	}()
	log.Println("shutting down")
	os.Exit(0)

	/*err := http.ListenAndServe(c.URL, handlers.CompressHandler(r))
	if err != nil {
		log.Fatal(err)
	}*/

}
