package main

import (
	"fmt"
	"net/http"
	"github.com/labstack/echo/v4"
)

// Echo framework server
func setupEchoServer() {
	e := echo.New()

	// Standard GET
	e.GET("/api/v3/users", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"users": []map[string]interface{}{{"id": 1, "name": "Charlie"}},
		})
	})

	// Standard POST
	e.POST("/api/v3/users", func(c echo.Context) error {
		var json map[string]interface{}
		c.Bind(&json)
		return c.JSON(http.StatusCreated, map[string]interface{}{
			"message": "User created",
			"data":    json,
		})
	})

	// Multiple methods on same endpoint
	e.GET("/api/v3/products", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"products": []map[string]interface{}{{"id": 1, "name": "Thing"}},
		})
	})
	e.POST("/api/v3/products", func(c echo.Context) error {
		var json map[string]interface{}
		c.Bind(&json)
		return c.JSON(http.StatusCreated, map[string]interface{}{
			"message": "Product created",
			"data":    json,
		})
	})
	e.PUT("/api/v3/products", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{"message": "Product updated"})
	})
	e.DELETE("/api/v3/products", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{"message": "Product deleted"})
	})

	// Non-standard method using Match
	e.Match([]string{"FUNKYTOWN"}, "/api/v3/funkytown", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{"message": "Echo Funkytown!"})
	})

	// Another non-standard method
	e.Match([]string{"DANCE"}, "/api/v3/dance", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{"message": "Echo Dance!"})
	})

	// Multiple non-standard methods
	e.Match([]string{"FUNKYTOWN", "PARTY"}, "/api/v3/custom", func(c echo.Context) error {
		method := c.Request().Method
		return c.JSON(http.StatusOK, map[string]interface{}{
			"method":    method,
			"framework": "echo",
		})
	})

	// Discouraged: Missing method check - using Any
	e.Any("/api/v3/bad/no-method-check", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "Accepts any method - bad practice!",
		})
	})

	// Discouraged: Missing error handling
	e.POST("/api/v3/bad/no-error-handling", func(c echo.Context) error {
		var json map[string]interface{}
		// No error check on Bind
		c.Bind(&json)
		return c.JSON(http.StatusOK, json)
	})

	// Discouraged: Missing input validation
	e.POST("/api/v3/bad/no-validation", func(c echo.Context) error {
		var json map[string]interface{}
		c.Bind(&json)
		// No validation
		return c.JSON(http.StatusOK, map[string]interface{}{"echo": json})
	})

	// Discouraged: Not returning error from handler
	e.GET("/api/v3/bad/no-error-return", func(c echo.Context) error {
		// Should return error but doesn't
		c.JSON(http.StatusOK, map[string]interface{}{"message": "Missing error return"})
		return nil // Should check for errors
	})

	fmt.Println("Echo server running on :8082")
	e.Start(":8082")
}

