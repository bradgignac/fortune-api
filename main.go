package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/bradgignac/fortune-api/api"
	"github.com/bradgignac/fortune-api/fortune"
)

var addr string
var db string
var timeout time.Duration

func init() {
	flag.StringVar(&addr, "addr", ":8000", "Address to bind to")
	flag.StringVar(&db, "db", "", "Path to fortune database")
	flag.DurationVar(&timeout, "timeout", time.Second*30, "HTTP request timeout")
}

func main() {
	flag.Parse()

	reader, err := os.Open(db)
	if err != nil {
		log.Fatal(err)
	}

	db, err := fortune.Parse(reader)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Initialized fortune database with %d fortunes", db.Count())

	api := api.NewHandler(db)
	server := &http.Server{
		Addr:         addr,
		ReadTimeout:  timeout,
		WriteTimeout: timeout,
		IdleTimeout:  timeout,
		Handler:      api,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	log.Printf("HTTP server is listening on %s", server.Addr)
	waitForShutdown(server)
}

func waitForShutdown(server *http.Server) {
	sigchan := make(chan os.Signal, 1)
	defer close(sigchan)

	signal.Notify(sigchan, os.Interrupt)
	received := <-sigchan

	log.Printf("Received %s. Shutting down...", received.String())
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	server.Shutdown(ctx)
}
