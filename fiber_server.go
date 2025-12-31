package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

// Fiber framework server
func setupFiberServer() {
	app := fiber.New()

	// Standard GET
	app.Get("/api/v4/users", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"users": []fiber.Map{{"id": 1, "name": "David"}},
		})
	})

	// Standard POST
	app.Post("/api/v4/users", func(c *fiber.Ctx) error {
		var json map[string]interface{}
		c.BodyParser(&json)
		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"message": "User created",
			"data":    json,
		})
	})

	// Multiple methods on same endpoint
	app.Get("/api/v4/products", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"products": []fiber.Map{{"id": 1, "name": "Widget"}},
		})
	})
	app.Post("/api/v4/products", func(c *fiber.Ctx) error {
		var json map[string]interface{}
		c.BodyParser(&json)
		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"message": "Product created",
			"data":    json,
		})
	})
	app.Put("/api/v4/products", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Product updated"})
	})
	app.Delete("/api/v4/products", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Product deleted"})
	})

	// Non-standard method using Add
	app.Add("FUNKYTOWN", "/api/v4/funkytown", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Fiber Funkytown!"})
	})

	// Another non-standard method
	app.Add("DANCE", "/api/v4/dance", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Fiber Dance!"})
	})

	// Multiple non-standard methods
	app.Add("FUNKYTOWN", "/api/v4/custom", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"method":    "funkytown",
			"framework": "fiber",
		})
	})
	app.Add("PARTY", "/api/v4/custom", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"method":    "party",
			"framework": "fiber",
		})
	})

	// Discouraged: Missing method check - using All
	app.All("/api/v4/bad/no-method-check", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Accepts any method - bad practice!",
		})
	})

	// Discouraged: Missing error handling
	app.Post("/api/v4/bad/no-error-handling", func(c *fiber.Ctx) error {
		var json map[string]interface{}
		// No error check on BodyParser
		c.BodyParser(&json)
		return c.JSON(json)
	})

	// Discouraged: Missing input validation
	app.Post("/api/v4/bad/no-validation", func(c *fiber.Ctx) error {
		var json map[string]interface{}
		c.BodyParser(&json)
		// No validation
		return c.JSON(fiber.Map{"echo": json})
	})

	// Discouraged: Not handling errors properly
	app.Get("/api/v4/bad/no-error-return", func(c *fiber.Ctx) error {
		// Should check for errors but doesn't
		c.JSON(fiber.Map{"message": "Missing error handling"})
		return nil
	})

	fmt.Println("Fiber server running on :8083")
	app.Listen(":8083")
}

