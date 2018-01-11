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
)

var addr string
var timeout time.Duration

func init() {
	flag.StringVar(&addr, "addr", ":8000", "Address to bind to")
	flag.DurationVar(&timeout, "timeout", time.Second*30, "HTTP request timeout")
}

func main() {
	flag.Parse()

	server := &http.Server{
		Addr:         addr,
		ReadTimeout:  timeout,
		WriteTimeout: timeout,
		IdleTimeout:  timeout,
		Handler:      &api.Handler{},
	}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	log.Printf("fortune-api is listening on %s!\n", server.Addr)
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
