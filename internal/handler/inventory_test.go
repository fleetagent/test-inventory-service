package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.GET("/health", HealthCheck)
	r.GET("/api/inventory", ListInventory)
	r.GET("/api/inventory/:sku", GetBySKU)
	r.POST("/api/inventory", CreateItem)
	r.PUT("/api/inventory/:sku", UpdateStock)
	r.DELETE("/api/inventory/:sku", DeleteItem)
	r.POST("/api/inventory/:sku/reserve", ReserveStock)
	r.POST("/api/inventory/:sku/release", ReleaseReservation)
	return r
}

func TestHealthCheck(t *testing.T) {
	r := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/health", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", w.Code)
	}
	var body map[string]string
	json.Unmarshal(w.Body.Bytes(), &body)
	if body["status"] != "ok" {
		t.Errorf("expected status ok, got %s", body["status"])
	}
}

func TestListInventory(t *testing.T) {
	r := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/inventory", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", w.Code)
	}
}

func TestGetBySKU(t *testing.T) {
	r := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/inventory/SKU-001", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", w.Code)
	}
	var body map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &body)
	if body["sku"] != "SKU-001" {
		t.Errorf("expected sku SKU-001, got %v", body["sku"])
	}
}

func TestCreateItem(t *testing.T) {
	r := setupRouter()
	item := InventoryItem{SKU: "NEW-001", Name: "Widget", Quantity: 50, Warehouse: "us-east"}
	body, _ := json.Marshal(item)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/inventory", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("expected 201, got %d", w.Code)
	}
	var result InventoryItem
	json.Unmarshal(w.Body.Bytes(), &result)
	if result.SKU != "NEW-001" {
		t.Errorf("expected SKU NEW-001, got %s", result.SKU)
	}
	if result.Quantity != 50 {
		t.Errorf("expected quantity 50, got %d", result.Quantity)
	}
}

func TestCreateItemBadJSON(t *testing.T) {
	r := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/inventory", bytes.NewBuffer([]byte("invalid")))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("expected 400, got %d", w.Code)
	}
}

func TestDeleteItem(t *testing.T) {
	r := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/api/inventory/SKU-001", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusNoContent {
		t.Errorf("expected 204, got %d", w.Code)
	}
}

func TestReserveStock(t *testing.T) {
	r := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/inventory/SKU-001/reserve", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", w.Code)
	}
	var body map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &body)
	if body["reserved"] != true {
		t.Errorf("expected reserved true")
	}
}

func TestReleaseReservation(t *testing.T) {
	r := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/inventory/SKU-001/release", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", w.Code)
	}
	var body map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &body)
	if body["released"] != true {
		t.Errorf("expected released true")
	}
}
