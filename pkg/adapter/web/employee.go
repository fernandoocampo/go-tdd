package web

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/fernandoocampo/go-tdd/pkg/employeeapp"
	"github.com/fernandoocampo/go-tdd/pkg/port"
)

// employeeRestHandler contains handler logic to create an employee.
type employeeRestHandler struct {
	service employeeapp.EmployeeService
}

// NewEmployeeRestHandler instance of a basic implementation of employee rest handler
func NewEmployeeRestHandler(employeeService employeeapp.EmployeeService) RestHandler {
	return employeeRestHandler{
		service: employeeService,
	}
}

// Create creates a new record
func (p employeeRestHandler) Create(w http.ResponseWriter, r *http.Request) {
	// context constraint
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	var employee port.NewEmployee
	// close the body buffer at the end of the function
	defer r.Body.Close()
	// Create the decoder for employee regarding to the body request
	decoder := json.NewDecoder(r.Body)
	// Get all the data of the request and map to employee struct
	// if error we response with error message
	if err := decoder.Decode(&employee); err != nil {
		RespondRestWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	newemployee := employee.AsEmployee()
	_, err := p.service.Add(ctx, *newemployee)
	if err != nil {
		RespondRestWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	RespondRestWithJSON(w, http.StatusOK, "created!")

}
