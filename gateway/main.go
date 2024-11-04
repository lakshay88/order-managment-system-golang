package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/go-chi/chi/v5"

	"github.com/lakshay88/order-managment-service-golang/commons/discovery"
	"github.com/lakshay88/order-managment-service-golang/gateway/routers"
)

var etcdClient *discovery.EtcdClient

func init() {
	endpoints := []string{"localhost:2379"}
	dialTimeout := 5 * time.Second
	etcdClient := discovery.NewEtcdClient(endpoints, dialTimeout)

	// Register the service (example)
	err := etcdClient.RegisterService(context.Background(), "gateway", "localhost:8080", 60)
	if err != nil {
		log.Fatalf("Failed to register service: %v", err)
	}
}

func main() {
	log.Println("Starting the gateway Service...")
	defer etcdClient.Close()

	// Initialize the Gin engine
	r := chi.NewRouter()
	// r.Use(middleware.Logger) // Logs requests
	// r.Use(middleware.Recoverer) // Recovers from panics
	apiRouter := routers.NewRouter()

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		apiRouter.RegisterRoutes(r)
	}()

	wg.Wait()

	// Set up the HTTP server with timeouts
	port := 8080
	srv := &http.Server{
		Addr:         fmt.Sprintf("0.0.0.0:%d", port),
		Handler:      r,                // the Gin handler
		ReadTimeout:  2 * time.Minute,  // Maximum duration for reading the request
		WriteTimeout: 0,                // Maximum duration before timing out writes (0 means no timeout)
		IdleTimeout:  60 * time.Second, // Maximum amount of time to keep an idle connection
	}

	// Start the server
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Failed to start server: %v", err)
	}
}
