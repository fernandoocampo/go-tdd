package port

import (
	"github.com/fernandoocampo/go-tdd/pkg/domain"
	"github.com/google/uuid"
)

// NewEmployeeName contains data related to the name of a employee
type NewEmployeeName struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

// NewEmployee contains data for a new employee
type NewEmployee struct {
	Name         NewEmployeeName `json:"names"`
	Salary       int64           `json:"salary"`
	JobCode      string          `json:"jobid"`
	JobName      string          `json:"job"`
	JobArea      string          `json:"jobarea"`
	JobDeparment string          `json:"jobdepartment"`
	Street       string          `json:"street"`
	City         string          `json:"city"`
	State        string          `json:"state"`
	ZipCode      string          `json:"zip"`
}

// AsEmployee converts the receiver NewEmployee to employee
func (ne *NewEmployee) AsEmployee() *domain.Employee {
	return &domain.Employee{
		ID:        uuid.New().String(),
		FirstName: ne.Name.FirstName,
		LastName:  ne.Name.LastName,
		Salary:    ne.Salary,
		Address: &domain.Address{
			Street:  ne.Street,
			City:    ne.City,
			State:   ne.State,
			ZipCode: ne.ZipCode,
		},
		Job: &domain.Job{
			Code:      ne.JobCode,
			Name:      ne.JobName,
			Area:      ne.JobArea,
			Deparment: ne.JobDeparment,
		},
	}
}
