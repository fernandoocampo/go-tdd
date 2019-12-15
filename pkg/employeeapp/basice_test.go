package employeeapp_test

import (
	"context"
	"testing"

	"github.com/fernandoocampo/go-tdd/pkg/domain"
	"github.com/fernandoocampo/go-tdd/pkg/employeeapp"
)

func TestAddEmployee(t *testing.T) {
	// GIVEN
	ctx := context.TODO()
	newRepo := aEmployeeRepoMock()
	newEmployee := aNewEmployeeDaenerys()
	basicService := employeeapp.NewBasicEmployeeSerive(newRepo)
	// WHEN
	got := basicService.Add(ctx, newEmployee)
	// THEN
	if got != nil {
		t.Errorf("no error was expected but got: %q", got)
	}
}

func aNewEmployeeDaenerys() domain.Employee {
	return domain.Employee{
		FirstName: "Daenerys",
		LastName:  "Targaryen",
		Salary:    25000,
		Address: &domain.Address{
			Street:  "Queen Street",
			City:    "King's Landing",
			State:   "Capital",
			ZipCode: "00021",
		},
		Job: &domain.Job{
			Code:      "100",
			Name:      "Manager",
			Area:      "Monarchy",
			Deparment: "Castle",
		},
	}
}

func aEmployeeRepoMock() domain.EmployeeRepository {
	t := struct{ ProtoRepo }{}

	t.data = make(map[string]*domain.Employee)
	t.queryByID = func(ctx context.Context, ID string) (*domain.Employee, error) {
		return t.data[ID], nil
	}
	t.save = func(ctx context.Context, newemployee domain.Employee) error {
		t.data[newemployee.ID] = &newemployee
		return nil
	}
	return t
}

type ProtoRepo struct {
	data      map[string]*domain.Employee
	queryByID func(ctx context.Context, ID string) (*domain.Employee, error)
	save      func(ctx context.Context, newemployee domain.Employee) error
}

func (t ProtoRepo) QueryByID(ctx context.Context, ID string) (*domain.Employee, error) {
	return t.queryByID(ctx, ID)
}
func (t ProtoRepo) Save(ctx context.Context, newemployee domain.Employee) error {
	return t.save(ctx, newemployee)
}
