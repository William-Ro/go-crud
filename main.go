package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/William-Ro/go-crud/api/routes"
	"github.com/William-Ro/go-crud/pkg/employee"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func main() {
	db, cancel, err := databaseConnection()
	if err != nil {
		log.Fatal("Database Connection Error $s", err)
	}
	fmt.Println("Database connection success!")
	employeeCollection := db.Collection("employees")
	employeeRepo := employee.NewRepo(employeeCollection)
	employeeService := employee.NewService(employeeRepo)

	app := fiber.New()
	app.Use(cors.New())
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Send([]byte("Welcome to the Go CRUD API"))
	})
	api := app.Group("/api")
	routes.EmployeeRouter(api, employeeService)
	defer cancel()
	log.Fatal(app.Listen(":8080"))
}

func databaseConnection() (*mongo.Database, context.CancelFunc, error) {

	// Load the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get the MongoDB URI from the environment variables
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environment variable. Example\n\t MONGODB_URI=mongodb://admin:admin@localhost:27017 See\n\t https://docs.mongodb.com/drivers/go/current/usage-examples/")
	}

	// Set client options
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	// Connect to MongoDB
	client, err := mongo.Connect(opts)
	if err != nil {
		return nil, nil, err
	}

	// Set a timeout for the connection
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// Ping the deployment to ensure you have a valid connection
	var result map[string]interface{}
	if err := client.Database("admin").RunCommand(ctx, map[string]interface{}{"ping": 1}).Decode(&result); err != nil {
		cancel()
		return nil, nil, err
	}

	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
	// Return the database instance
	db := client.Database("employee_management")
	return db, cancel, nil
}
