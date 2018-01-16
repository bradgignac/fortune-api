package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/bradgignac/fortune-api/api"
	"github.com/bradgignac/fortune-api/fortune"
)

var addr string
var timeout time.Duration

func init() {
	flag.StringVar(&addr, "addr", ":8000", "Address to bind to")
	flag.DurationVar(&timeout, "timeout", time.Second*30, "HTTP request timeout")
}

func main() {
	flag.Parse()

	data := `"Failure is the opportunity to begin again more intelligently."
  ~Henry Ford
%
"one more"
  ~Unknown`
	db, err := fortune.Parse(strings.NewReader(data))
	if err != nil {
		log.Fatal(err)
	}

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
