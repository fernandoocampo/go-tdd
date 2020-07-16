package employeeapp

import (
	"context"

	"github.com/fernandoocampo/go-tdd/pkg/domain"
	"github.com/google/uuid"
)

// basicEmployeeService implements employee service
type basicEmployeeService struct {
	repository domain.EmployeeRepository
}

// NewBasicEmployeeService creates a employee service using basic employee service impl.
func NewBasicEmployeeService(newRepository domain.EmployeeRepository) Service {
	return &basicEmployeeService{
		repository: newRepository,
	}
}

// Add validates and save a new employee.
func (b *basicEmployeeService) Add(ctx context.Context, newemployee domain.Employee) (string, error) {
	newemployee.ID = uuid.New().String()
	err := b.repository.Save(ctx, newemployee)
	return newemployee.ID, err
}
