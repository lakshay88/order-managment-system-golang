package discovery

import (
	"context"
	"log"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

type EtcdClient struct {
	client *clientv3.Client
}

func NewEtcdClient(endpoints []string, dialTimeout time.Duration) *EtcdClient {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: dialTimeout,
	})
	if err != nil {
		log.Fatalf("Failed to connect to etcd: %v", err)
	}
	return &EtcdClient{client: cli}
}

// RegisterService registers a service with a specified name and address
func (e *EtcdClient) RegisterService(ctx context.Context, serviceName, serviceAddr string, ttl int) error {
	// Create a lease with the specified TTL
	leaseResp, err := e.client.Grant(ctx, int64(ttl))
	if err != nil {
		return err
	}

	// Register the service with the lease
	_, err = e.client.Put(ctx, serviceName, serviceAddr, clientv3.WithLease(leaseResp.ID))
	if err != nil {
		return err
	}
	log.Printf("Service %s registered at %s with TTL of %d seconds", serviceName, serviceAddr, ttl)

	// Keep the lease alive in a separate goroutine if necessary
	go e.keepLeaseAlive(ctx, leaseResp.ID)

	return nil
}

func (e *EtcdClient) keepLeaseAlive(ctx context.Context, leaseID clientv3.LeaseID) {
	ch, err := e.client.KeepAlive(ctx, leaseID)
	if err != nil {
		log.Printf("Failed to keep lease alive: %v", err)
		return
	}
	for {
		select {
		case kaResp := <-ch:
			if kaResp == nil {
				log.Printf("Lease %v expired", leaseID)
				return
			}
			log.Printf("Lease %v renewed", kaResp.ID)
		case <-ctx.Done():
			log.Println("Stopping lease renewal")
			return
		}
	}
}

// DeregisterService deregisters a service by name
func (e *EtcdClient) DeregisterService(ctx context.Context, serviceName string) error {
	_, err := e.client.Delete(ctx, serviceName)
	if err != nil {
		return err
	}
	log.Printf("Service %s deregistered", serviceName)
	return nil
}

// Close closes the etcd client connection
func (e *EtcdClient) Close() {
	if err := e.client.Close(); err != nil {
		log.Printf("Failed to close etcd client: %v", err)
	} else {
		log.Println("Etcd client closed successfully")
	}
}
