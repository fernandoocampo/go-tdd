package domain

import "context"

// EmployeeRepository defines behavior for employee persistence.
type EmployeeRepository interface {
	QueryByID(ctx context.Context, ID string) (*Employee, error)
	Save(ctx context.Context, newemployee Employee) error
}
