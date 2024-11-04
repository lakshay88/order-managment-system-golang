package api

import (
	"net/http"

	"github.com/go-chi/chi"
)

type OrderRoutes struct{}

func (o OrderRoutes) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/orders", o.GetOrdersItems) // Get all orders
	// r.Post("/orders", r.CreateOrder) // Create a new order
	// router.GET("/orders/:id", r.GetOrderByID)   // Get an order by ID
	// router.PUT("/orders/:id", r.UpdateOrder)    // Update an existing order
	// router.DELETE("/orders/:id", r.DeleteOrder) // Delete an order

	return r
}

func (o OrderRoutes) GetOrdersItems(w http.ResponseWriter, req *http.Request) {
	// Implementation here
}
