package employeeapp

import (
	"context"

	"github.com/fernandoocampo/go-tdd/pkg/domain"
)

// Service defines employee service behavior
type Service interface {
	// Add validates and save a new employee. Returns new employee id and an error
	// if something is wrong.
	Add(ctx context.Context, newemployee domain.Employee) (string, error)
}
