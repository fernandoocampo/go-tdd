package web_test

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"testing"

	"github.com/fernandoocampo/go-tdd/pkg/adapter/web"
	"github.com/fernandoocampo/go-tdd/pkg/domain"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/mock"
)

func TestCreateAEmployee(t *testing.T) {
	service := new(testifyMock)
	service.On("Add", mock.AnythingOfType("*context.timerCtx"), mock.AnythingOfType("Employee")).Return("1", nil)
	employeehandler := web.NewEmployeeRestHandler(service)

	strjson := string(loadTestFixturesObjects(t, "the_hound_employee.json"))
	req, errreq := http.NewRequest("POST", "/employees", bytes.NewBuffer([]byte(strjson)))

	if errreq != nil {
		t.Fatal(errreq)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	r := mux.NewRouter()
	r.HandleFunc("/employees", employeehandler.Create).Methods("POST")

	// When client consumes a rest api.
	r.ServeHTTP(rr, req)

	// Then we check the result of the employee found.
	service.AssertExpectations(t)
	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

// loadTestFixturesObjects helps to load files used as fixtures on tests
func loadTestFixturesObjects(t *testing.T, name string) []byte {
	t.Helper()
	path := filepath.Join("testdata", name)
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	return bytes
}

type testifyMock struct {
	mock.Mock // just for academic purposes
}

func (t *testifyMock) Add(ctx context.Context, newemployee domain.Employee) (string, error) {
	args := t.Called(ctx, newemployee)
	return args.String(0), args.Error(1)
}
