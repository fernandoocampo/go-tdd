package employeeapp

import (
	"context"

	"github.com/fernandoocampo/go-tdd/pkg/domain"
)

// EmployeeService defines employee service behavior
type EmployeeService interface {
	// Add validates and save a new employee.
	Add(ctx context.Context, newemployee domain.Employee) error
}
