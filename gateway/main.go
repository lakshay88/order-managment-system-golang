package main

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
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
	r := gin.Default() // Initialize your router
	apiRouter := routers.NewRouter()

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		apiRouter.RegisterRoutes(r)
	}()

	wg.Wait()

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
