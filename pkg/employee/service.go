package employee

import (
	"github.com/William-Ro/go-crud/api/presenter"
	"github.com/William-Ro/go-crud/pkg/entities"
)

// Service is an interface from which our api module can access our repository of all our models
type Service interface {
	InsertEmployee(employee *entities.Employee) (*entities.Employee, error)
	FetchEmployees() (*[]presenter.Employee, error)
	UpdateEmployee(employee *entities.Employee) (*entities.Employee, error)
	RemoveEmployee(ID string) error
}

type service struct {
	repository Repository
}

// NewService is used to create a single instance of the service
func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

// InsertEmployee is a service layer that helps insert employees
func (s *service) InsertEmployee(employee *entities.Employee) (*entities.Employee, error) {
	return s.repository.CreateEmployee(employee)
}

// FetchEmployees is a service layer that helps fetch all employees
func (s *service) FetchEmployees() (*[]presenter.Employee, error) {
	return s.repository.ReadEmployee()
}

// UpdateEmployee is a service layer that helps update employees
func (s *service) UpdateEmployee(employee *entities.Employee) (*entities.Employee, error) {
	return s.repository.UpdateEmployee(employee)
}

// RemoveEmployee is a service layer that helps remove employees
func (s *service) RemoveEmployee(ID string) error {
	return s.repository.DeleteEmployee(ID)
}
