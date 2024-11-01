package routers

import "github.com/gin-gonic/gin"

type Routers struct{}

func NewRouter() *Routers {
	return &Routers{}
}

func (r *Routers) RegisterRoutes(router *gin.Engine) {
	router.GET("/orders", r.GetOrders)          // Get all orders
	router.POST("/orders", r.CreateOrder)       // Create a new order
	router.GET("/orders/:id", r.GetOrderByID)   // Get an order by ID
	router.PUT("/orders/:id", r.UpdateOrder)    // Update an existing order
	router.DELETE("/orders/:id", r.DeleteOrder) // Delete an order
}
