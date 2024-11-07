package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"reddit-api/internal/monitor"
	"reddit-api/internal/server"
	"reddit-api/internal/store"
	"sync"
	"time"
)

var httpPort = flag.String("port", "80", "http listen port")
var host = flag.String("host", "127.0.0.1", "host name")
var clientId = flag.String("id", "", "client id")
var clientSecret = flag.String("secret", "", "client id")
var subreddit = flag.String("subreddit", "/r/programming", "subreddit to crawl")

func main() {
	ctx := context.Background()
	if err := run(ctx); err != nil {
		fmt.Fprintf(os.Stderr, "%\n", err)
		os.Exit(1)
	}
}

func run(ctx context.Context) error {
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	flag.Parse()

	var us store.UserStore
	us.Users = make(map[string]store.User)
	var ps store.PostStore
	ps.Posts = make(map[string]store.Post)

	srv := server.CreateServer(&us, &ps)

	httpServer := &http.Server{
		Addr:    net.JoinHostPort(*host, *httpPort),
		Handler: srv,
	}

	go func() {
		fmt.Printf("listening on %s\n", httpServer.Addr)
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Fprintf(os.Stderr, "error listening and serving: %s\n", err)
		}
	}()

	monitor.StartMonitor(&us, &ps, *clientId, *clientSecret, *subreddit)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		<-ctx.Done()
		shutdownCtx := context.Background()
		shutdownCtx, cancel := context.WithTimeout(shutdownCtx, 10*time.Second)
		defer cancel()
		if err := httpServer.Shutdown(shutdownCtx); err != nil {
			fmt.Fprintf(os.Stderr, "failed to shut down http server: %s\n", err)
		}
	}()
	wg.Wait()
	return nil
}
