package main

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)

// Gin framework server
func setupGinServer() {
	r := gin.Default()

	// Standard GET
	r.GET("/api/v2/users", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"users": []gin.H{{"id": 1, "name": "Bob"}}})
	})

	// Standard POST
	r.POST("/api/v2/users", func(c *gin.Context) {
		var json map[string]interface{}
		c.BindJSON(&json)
		c.JSON(http.StatusCreated, gin.H{"message": "User created", "data": json})
	})

	// Multiple methods on same endpoint
	r.GET("/api/v2/products", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"products": []gin.H{{"id": 1, "name": "Gadget"}}})
	})
	r.POST("/api/v2/products", func(c *gin.Context) {
		var json map[string]interface{}
		c.BindJSON(&json)
		c.JSON(http.StatusCreated, gin.H{"message": "Product created", "data": json})
	})
	r.PUT("/api/v2/products", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Product updated"})
	})
	r.DELETE("/api/v2/products", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Product deleted"})
	})

	// Non-standard method using Handle
	r.Handle("FUNKYTOWN", "/api/v2/funkytown", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Gin Funkytown!"})
	})

	// Another non-standard method
	r.Handle("DANCE", "/api/v2/dance", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Gin Dance!"})
	})

	// Multiple non-standard methods
	r.Handle("FUNKYTOWN", "/api/v2/custom", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"method": "funkytown", "framework": "gin"})
	})
	r.Handle("PARTY", "/api/v2/custom", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"method": "party", "framework": "gin"})
	})

	// Discouraged: Missing method validation - using Any
	r.Any("/api/v2/bad/no-method-check", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Accepts any method - bad practice!"})
	})

	// Discouraged: Missing error handling
	r.POST("/api/v2/bad/no-error-handling", func(c *gin.Context) {
		var json map[string]interface{}
		// No error check on BindJSON
		c.BindJSON(&json)
		c.JSON(http.StatusOK, json)
	})

	// Discouraged: Missing input validation
	r.POST("/api/v2/bad/no-validation", func(c *gin.Context) {
		var json map[string]interface{}
		c.BindJSON(&json)
		// No validation
		c.JSON(http.StatusOK, gin.H{"echo": json})
	})

	// Discouraged: Using deprecated method
	r.HandleMethodNotAllowed = true
	r.NoMethod(func(c *gin.Context) {
		// But then not handling it properly
		c.JSON(http.StatusOK, gin.H{"message": "Method not allowed but returning 200"})
	})

	fmt.Println("Gin server running on :8081")
	r.Run(":8081")
}

