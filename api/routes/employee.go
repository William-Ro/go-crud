package routes

import (
	"github.com/William-Ro/go-crud/api/handlers"
	"github.com/William-Ro/go-crud/pkg/employee"
	"github.com/gofiber/fiber/v2"
)

// EmployeeRouter is the Router for GoFiber App
func EmployeeRouter(app fiber.Router, service employee.Service) {
	app.Get("/employees", handlers.GetEmployees(service))
	app.Post("/employees", handlers.AddEmployee(service))
	app.Put("/employees", handlers.UpdateEmployee(service))
	app.Delete("/employees", handlers.RemoveEmployee(service))
}
