package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type InventoryItem struct {
	SKU       string `json:"sku"`
	Name      string `json:"name"`
	Quantity  int    `json:"quantity"`
	Reserved  int    `json:"reserved"`
	Warehouse string `json:"warehouse"`
}

func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func ListInventory(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"items": []InventoryItem{}, "total": 0})
}

func GetBySKU(c *gin.Context) {
	sku := c.Param("sku")
	c.JSON(http.StatusOK, gin.H{"sku": sku, "quantity": 0})
}

func CreateItem(c *gin.Context) {
	var item InventoryItem
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, item)
}

func UpdateStock(c *gin.Context) {
	sku := c.Param("sku")
	c.JSON(http.StatusOK, gin.H{"sku": sku, "updated": true})
}

func DeleteItem(c *gin.Context) {
	c.Status(http.StatusNoContent)
}

func ReserveStock(c *gin.Context) {
	sku := c.Param("sku")
	c.JSON(http.StatusOK, gin.H{"sku": sku, "reserved": true})
}

func ReleaseReservation(c *gin.Context) {
	sku := c.Param("sku")
	c.JSON(http.StatusOK, gin.H{"sku": sku, "released": true})
}
