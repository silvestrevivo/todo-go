package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v3"
)

type Todo struct {
	ID        int    `json:"id"`
	Completed bool   `json:"completed"`
	Body      string `json:"body"`
}

func main() {
	fmt.Println("Hello World!!!!!!!!!!!")

	// Initialize a new Fiber app
	app := fiber.New()

	todos := []Todo{}

	// Define a route for the GET method on the root path '/'
	app.Get("/", func(c fiber.Ctx) error {
		// Send a string response to the client
		return c.Status(200).JSON(fiber.Map{"msg": "hello world"})
	})

	app.Get("/api/todos", func(c fiber.Ctx) error {
		return c.Status(200).JSON(todos)
	})

	app.Get("/api/todos/:id", func(c fiber.Ctx) error {
		id := c.Params("id")

		for i, t := range todos {
			idInt, err := strconv.Atoi(id)
			if err != nil {
				fmt.Println("Error converting string to int:", err)
				break
			}

			if t.ID == idInt {
				return c.Status(201).JSON(todos[i])
			}
		}

		return c.Status(200).JSON(todos)
	})

	app.Post("/api/todos", func(c fiber.Ctx) error {
		todo := new(Todo)

		if err := c.Bind().Body(todo); err != nil {
			return err
		}

		if todo.Body == "" {
			return c.Status(400).JSON(fiber.Map{"error": "body is required"})
		}

		todo.ID = len(todos) + 1
		todos = append(todos, *todo)

		return c.Status(201).JSON(todos)
	})

	app.Patch("/api/todos/:id", func(c fiber.Ctx) error {
		id := c.Params("id")
		todo := new(Todo)

		if err := c.Bind().Body(todo); err != nil {
			return err
		}

		for i, t := range todos {
			idInt, err := strconv.Atoi(id)
			if err != nil {
				fmt.Println("Error converting string to int:", err)
				break
			}

			if t.ID == idInt {
				todoAdjusted := new(Todo)
				todoAdjusted.ID = t.ID
				todoAdjusted.Body = todo.Body
				todoAdjusted.Completed = todo.Completed
				todos[i] = *todoAdjusted
				return c.Status(201).JSON(todos)
			}
		}

		return c.Status(400).JSON(fiber.Map{"error": "id or user not found"})

	})

	app.Delete("/api/todos/:id", func(c fiber.Ctx) error {
		id := c.Params("id")

		for i, t := range todos {
			idInt, err := strconv.Atoi(id)
			if err != nil {
				fmt.Println("Error converting string to int:", err)
				break
			}

			if t.ID == idInt {
				todos = append(todos[:i], todos[i+1:]...)
				return c.Status(201).JSON(todos)
			}
		}

		return c.Status(400).JSON(fiber.Map{"error": "id or user not found"})
	})

	log.Fatal(app.Listen(":8080"))
}
