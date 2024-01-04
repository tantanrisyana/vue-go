// main.go
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Endpoint untuk mendapatkan data DataGrid
	router.GET("/api/dataGrid", func(c *gin.Context) {
		// Implementasikan logika untuk mengambil data DataGrid dari database atau sumber lainnya
		// Contoh sederhana:
		data := []map[string]interface{}{
			{"id": 1, "name": "Item 1", "description": "Description 1"},
			{"id": 2, "name": "Item 2", "description": "Description 2"},
			// ...
		}
		c.JSON(http.StatusOK, data)
	})

	// Endpoint untuk mendapatkan data Tree
	router.GET("/api/tree", func(c *gin.Context) {
		// Implementasikan logika untuk mengambil data Tree dari database atau sumber lainnya
		// Contoh sederhana:
		treeData := []map[string]interface{}{
			{"id": 1, "label": "Node 1", "children": []map[string]interface{}{{"id": 2, "label": "Node 1.1"}}},
			{"id": 3, "label": "Node 2"},
			// ...
		}
		c.JSON(http.StatusOK, treeData)
	})

	router.Run(":8080")
}
