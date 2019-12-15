package port_test

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/fernandoocampo/go-tdd/pkg/domain"
	"github.com/stretchr/testify/assert"

	"github.com/fernandoocampo/go-tdd/pkg/port"
)

func TestNewEmployeeAsEmployee(t *testing.T) {
	t.Run("memory_fixture", func(t *testing.T) {
		// t.Parallel()
		// GIVEN
		newEmployee := aNewEmployeeDaenerys()
		want := &domain.Employee{
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
		// WHEN
		got := newEmployee.AsEmployee()
		want.ID = got.ID
		// time.Sleep(1 * time.Second)
		// THEN
		assert.Equal(t, want, got)
	})
	t.Run("file_fixture_hound", func(t *testing.T) {
		// t.Parallel()
		// GIVEN
		newEmployee := buildEmployeeFromFileHelper(t, "the_hound_employee.json")
		want := &domain.Employee{
			FirstName: "Sandor",
			LastName:  "Clegane",
			Salary:    100,
			Address: &domain.Address{
				Street:  "House Clegane",
				City:    "King's Landing",
				State:   "Capital",
				ZipCode: "20021",
			},
			Job: &domain.Job{
				Code:      "234",
				Name:      "Bodyguard",
				Area:      "King's guard",
				Deparment: "Castle",
			},
		}
		// WHEN
		got := newEmployee.AsEmployee()
		want.ID = got.ID
		// time.Sleep(1 * time.Second)
		// THEN
		assert.Equal(t, want, got)
	})
	t.Run("file_fixture_mountain", func(t *testing.T) {
		// t.Parallel()
		// GIVEN
		newEmployee := buildEmployeeFromFileHelper(t, "the_mountain_employee.json")
		want := &domain.Employee{
			FirstName: "Gregor",
			LastName:  "Clegane",
			Salary:    200,
			Address: &domain.Address{
				Street:  "House Clegane",
				City:    "King's Landing",
				State:   "Capital",
				ZipCode: "20022",
			},
			Job: &domain.Job{
				Code:      "234",
				Name:      "Bodyguard",
				Area:      "King's guard",
				Deparment: "Castle",
			},
		}
		// WHEN
		got := newEmployee.AsEmployee()
		want.ID = got.ID
		// time.Sleep(1 * time.Second)
		// THEN
		assert.Equal(t, want, got)
	})
}

func aNewEmployeeDaenerys() *port.NewEmployee {
	return &port.NewEmployee{
		Name: port.NewEmployeeName{
			FirstName: "Daenerys",
			LastName:  "Targaryen",
		},
		Salary:       25000,
		JobCode:      "100",
		JobName:      "Manager",
		JobArea:      "Monarchy",
		JobDeparment: "Castle",
		Street:       "Queen Street",
		City:         "King's Landing",
		State:        "Capital",
		ZipCode:      "00021",
	}
}

// buildEmployeeFromFileHelper creates an employee request from a given json file
func buildEmployeeFromFileHelper(t *testing.T, filename string) *port.NewEmployee {
	t.Helper()
	bytesOrderWithTaxExempt := loadTestFixturesObjects(t, filename)
	request := &port.NewEmployee{}
	if err := json.Unmarshal(bytesOrderWithTaxExempt, request); err != nil {
		t.Error("error in unmarshalling data to request")
		t.Errorf("%v", err)
		t.FailNow()
	}
	return request
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
