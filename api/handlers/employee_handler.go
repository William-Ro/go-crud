package handlers

import (
	"errors"
	"net/http"

	"github.com/William-Ro/go-crud/api/presenter"
	"github.com/William-Ro/go-crud/pkg/employee"
	"github.com/William-Ro/go-crud/pkg/entities"
	"github.com/gofiber/fiber/v2"
)

// AddEmployee is handler/controller which creates employees
func AddEmployee(service employee.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody entities.Employee
		err := c.BodyParser(&requestBody)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.EmployeeErrorResponse(err))
		}
		if requestBody.First_Name == "" || requestBody.Last_Name == "" || requestBody.Email == "" {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.EmployeeErrorResponse(errors.New(
				"please specify all the fields")))
		}
		result, err := service.InsertEmployee(&requestBody)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.EmployeeErrorResponse(err))
		}
		return c.JSON(presenter.EmployeeSuccessResponse(result))
	}
}

// UpdateEmployee is handler/controller which updates data of Employees
func UpdateEmployee(service employee.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody entities.Employee
		err := c.BodyParser(&requestBody)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.EmployeeErrorResponse(err))
		}
		result, err := service.UpdateEmployee(&requestBody)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.EmployeeErrorResponse(err))
		}
		return c.JSON(presenter.EmployeeSuccessResponse(result))
	}
}

// RemoveEmployee is handler/controller which removes Employees
func RemoveEmployee(service employee.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody entities.DeleteRequest
		err := c.BodyParser(&requestBody)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.EmployeeErrorResponse(err))
		}
		employeeID := requestBody.ID
		err = service.RemoveEmployee(employeeID)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.EmployeeErrorResponse(err))
		}
		return c.JSON(&fiber.Map{
			"status": true,
			"data":   "updated successfully",
			"err":    nil,
		})
	}
}

// GetEmployees is handler/controller which lists all Employees
func GetEmployees(service employee.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		fetched, err := service.FetchEmployees()
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.EmployeeErrorResponse(err))
		}
		return c.JSON(presenter.EmployeesSuccessResponse(fetched))
	}
}
