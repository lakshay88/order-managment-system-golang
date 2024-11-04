package routers

import (
	"github.com/go-chi/chi/v5"
	"github.com/lakshay88/order-managment-service-golang/gateway/routers/api"
)

type Routers struct{}

func NewRouter() *Routers {
	return &Routers{}
}

func (r *Routers) RegisterRoutes(router *chi.Mux) {

	// Kitchen routs -
	// router.Mount("/api/", KitchenRoutes{}.Routes())

	router.Mount("/api", api.OrderRoutes{}.Routes())

	// router.Mount("/api", OrderRoutes{}.Routes())

}

type MenuRoutes struct{}

func (o MenuRoutes) Routes() chi.Router {
	r := chi.NewRouter()

	// r.Get("/orders", r.GetOrders)    // Get all orders
	// r.Post("/orders", r.CreateOrder) // Create a new order
	// router.GET("/orders/:id", r.GetOrderByID)   // Get an order by ID
	// router.PUT("/orders/:id", r.UpdateOrder)    // Update an existing order
	// router.DELETE("/orders/:id", r.DeleteOrder) // Delete an order

	return r
}
