package employeeapp_test

import (
	"context"
	"testing"

	"github.com/fernandoocampo/go-tdd/pkg/domain"
	"github.com/fernandoocampo/go-tdd/pkg/employeeapp"
	"github.com/stretchr/testify/mock"
)

func TestAddEmployee(t *testing.T) {
	t.Run("with_local_mock", func(t *testing.T) {
		// GIVEN
		ctx := context.TODO()
		newRepo := aEmployeeRepoMock()
		newEmployee := aNewEmployeeDaenerys()
		basicService := employeeapp.NewBasicEmployeeSerive(newRepo)
		// WHEN
		got, err := basicService.Add(ctx, newEmployee)
		// THEN
		if err != nil {
			t.Errorf("no error was expected but got: %q", err)
		}

		storedEmployee, queryErr := newRepo.QueryByID(ctx, got)
		if queryErr != nil {
			t.Errorf("no error was expected but got: %q", queryErr)
		}
		if storedEmployee == nil {
			t.Errorf("new employee Daenerys was not stored on database")
		}
	})

	t.Run("with_testify_mock", func(t *testing.T) {
		// GIVEN
		ctx := context.TODO()
		newRepo := new(testifyMock)
		newEmployee := aNewEmployeeDaenerys()
		// setup expectations
		newRepo.On("Save", mock.AnythingOfType("*context.emptyCtx"), mock.AnythingOfType("Employee")).Return(nil)
		basicService := employeeapp.NewBasicEmployeeSerive(newRepo)
		// WHEN
		got, err := basicService.Add(ctx, newEmployee)
		// THEN
		// assert that the expectations were met
		newRepo.AssertExpectations(t)
		if err != nil {
			t.Errorf("no error was expected but got: %q", err)
		}

		if got == "" {
			t.Errorf("expected an ID with value but got empty ID")
		}
	})
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

// aEmployeeRepoMock creates a mock from employee repository interface.
func aEmployeeRepoMock() domain.EmployeeRepository {
	t := &protoRepo{}
	// t := &struct{ protoRepo }{}

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

// protoRepo mocks a employee repository.
type protoRepo struct {
	data      map[string]*domain.Employee
	queryByID func(ctx context.Context, ID string) (*domain.Employee, error)
	save      func(ctx context.Context, newemployee domain.Employee) error
}

func (t *protoRepo) QueryByID(ctx context.Context, ID string) (*domain.Employee, error) {
	return t.queryByID(ctx, ID)
}
func (t *protoRepo) Save(ctx context.Context, newemployee domain.Employee) error {
	return t.save(ctx, newemployee)
}

type testifyMock struct {
	mock.Mock // just for academic purposes
}

func (t *testifyMock) QueryByID(ctx context.Context, ID string) (*domain.Employee, error) {
	args := t.Called(ctx, ID)
	employee := args.Get(0)
	err := args.Error(1)

	if employee == nil {
		return nil, err
	}

	return employee.(*domain.Employee), err
}
func (t *testifyMock) Save(ctx context.Context, newemployee domain.Employee) error {
	args := t.Called(ctx, newemployee)
	return args.Error(0)
}
