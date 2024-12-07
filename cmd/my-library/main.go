package main

import (
	"context"
	"fmt"
	"github/SpaceBuckett/books-api/cmd/internal/config"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	fmt.Println("WELCOME TO YOUR PERSONAL LIBRARY")
	fmt.Println("--------------------------------")

	// TODO: LOAD CONFIG
	cfg := config.LoadConfig()
	// TODO: SETUP DATABASE
	// TODO: SETUP ROUTER
	router := http.NewServeMux()

	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Shahzad's Dump!"))
	})
	// TODO: SETUP HTTP SERVER
	server := http.Server{
		Addr:    cfg.Addr,
		Handler: router,
	}

	fmt.Println("SERVER STARTED...")
	// TODO: ADD GRACEFULL SHUTDOWNS
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal("FAILED TO START SERVER")
		}

	}()

	<-done

	slog.Info("SHUTTING DOWN THE SERVER!")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := server.Shutdown(ctx)
	if err != nil {
		slog.Error("FAILED TO SHUTDHOWN THE SERVERS", slog.String("ERROR: ", err.Error()))
	}

	slog.Info("SERVERS SHUTDOWN SUCCESSFUL")
	// TODO: ADD CLI COMMANDS
	// TODO: ADD UNIT TESTS
	// TODO: ADD DOCUMENTATION
	// TODO: ADD CI/CD PIPELINE
	// TODO: ADD LOGGING
	// TODO: ADD ERROR HANDLING
	// TODO: ADD OPTIMIZATION
	// TODO: ADD DATABASE SUPPORT
	// TODO: ADD AUTHENTICATION
	// TODO: ADD CACHING
	// TODO: ADD PERFORMANCE TESTING
	// TODO: ADD SECURITY
	// TODO: ADD CODE OF CONDUCT
	// TODO: ADD LICENSE
	// TODO: ADD VERSION
	// TODO: ADD SUPPORT
	// TODO: ADD TESTING
	// TODO: ADD RELEASE

}
