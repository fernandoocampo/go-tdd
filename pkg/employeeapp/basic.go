package employeeapp

import (
	"context"

	"github.com/fernandoocampo/go-tdd/pkg/domain"
)

// basicEmployeeService implements employee service
type basicEmployeeService struct {
	repository domain.EmployeeRepository
}

// NewBasicEmployeeSerive creates a employee service using basic employee service impl.
func NewBasicEmployeeSerive(newRepository domain.EmployeeRepository) EmployeeService {
	return &basicEmployeeService{
		repository: newRepository,
	}
}

// Add validates and save a new employee.
func (b *basicEmployeeService) Add(ctx context.Context, newemployee domain.Employee) error {
	return b.repository.Save(ctx, newemployee)
}
