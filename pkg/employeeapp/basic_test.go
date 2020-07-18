package employeeapp_test

import (
	"context"
	"sync"
	"testing"

	"github.com/fernandoocampo/go-tdd/pkg/domain"
	"github.com/fernandoocampo/go-tdd/pkg/employeeapp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAddEmployee(t *testing.T) {
	t.Run("with_local_mock", func(t *testing.T) {
		// GIVEN
		ctx := context.TODO()
		newRepo := aEmployeeRepoMock()
		newEmployee := aNewEmployeeDaenerys()
		basicService := employeeapp.NewBasicService(newRepo)
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
		basicService := employeeapp.NewBasicService(newRepo)
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

func TestEmailNotification(t *testing.T) {
	// GIVEN
	newEmployee := aNewEmployeeDaenerys()
	newEmployee.Email = "someone@somewhere.com"
	expectedMessage := domain.Message{
		Subject: "Hey",
		To:      "someone@somewhere.com",
		From:    "anybody@somewhere.com",
		Body:    "Good morning there",
	}
	ctx := context.TODO()
	emailNotifier := NewEmailNotifierMock()
	newRepo := aEmployeeRepoMock()
	basicService := employeeapp.NewBasicServiceWithNotifier(newRepo, emailNotifier)

	// WHEN
	emailNotifier.wg.Add(1)
	_, err := basicService.Add(ctx, newEmployee)
	emailNotifier.wg.Wait()
	// THEN
	if err != nil {
		t.Errorf("no error was expected but got: %q", err)
	}

	assert.Equal(t, expectedMessage, emailNotifier.messages[0])
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
	t := &struct{ protoRepo }{}

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

// emailNotifierMock is a domain.Notifier mock that track
// received messages
type emailNotifierMock struct {
	err      error
	messages []domain.Message
	wg       sync.WaitGroup
}

func NewEmailNotifierMock() *emailNotifierMock {
	return &emailNotifierMock{
		messages: make([]domain.Message, 0),
	}
}

func NewEmailNotifierMockWithError(err error) *emailNotifierMock {
	newEmailNotifierMock := NewEmailNotifierMock()
	newEmailNotifierMock.err = err
	return newEmailNotifierMock
}

// Notify adds the given message to the messages map contained in the mock.
func (e *emailNotifierMock) Notify(ctx context.Context, message domain.Message) error {
	defer e.wg.Done()
	if e.err != nil {
		return e.err
	}
	e.messages = append(e.messages, message)
	return nil
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
