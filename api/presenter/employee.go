package presenter

import (
	"github.com/William-Ro/go-crud/pkg/entities"
	"github.com/gofiber/fiber"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Employee is the presenter object which will be passed in the response by Handler
type Employee struct {
	ID         primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	First_Name string             `json:"first_name"`
	Last_Name  string             `json:"last_name"`
	Email      string             `json:"email"`
}

// EmployeeSuccessResponse is the singular SuccessResponse that will be passed in the response by
// Handler
func EmployeeSuccessResponse(data *entities.Employee) *fiber.Map {
	employee := Employee{
		ID:         data.ID,
		First_Name: data.First_Name,
		Last_Name:  data.Last_Name,
		Email:      data.Email,
	}
	return &fiber.Map{
		"status": true,
		"data":   employee,
		"error":  nil,
	}
}

// EmployeesSuccessResponse is the list SuccessResponse that will be passed in the response by Handler
func EmployeesSuccessResponse(data *[]Employee) *fiber.Map {
	return &fiber.Map{
		"status": true,
		"data":   data,
		"error":  nil,
	}
}

// EmployeeErrorResponse is the ErrorResponse that will be passed in the response by Handler
func EmployeeErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"status": false,
		"data":   "",
		"error":  err.Error(),
	}
}
