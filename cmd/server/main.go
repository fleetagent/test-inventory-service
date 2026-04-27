package main

import (
	"log"
	"os"

	"github.com/fleetagent/inventory-service/internal/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/health", handler.HealthCheck)
	r.GET("/api/inventory", handler.ListInventory)
	r.GET("/api/inventory/:sku", handler.GetBySKU)
	r.POST("/api/inventory", handler.CreateItem)
	r.PUT("/api/inventory/:sku", handler.UpdateStock)
	r.DELETE("/api/inventory/:sku", handler.DeleteItem)
	r.POST("/api/inventory/:sku/reserve", handler.ReserveStock)
	r.POST("/api/inventory/:sku/release", handler.ReleaseReservation)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3003"
	}
	log.Printf("Inventory service starting on :%s", port)
	log.Fatal(r.Run(":" + port))
}
