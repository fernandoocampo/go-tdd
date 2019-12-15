package domain

// Address contains address data.
type Address struct {
	Street  string
	City    string
	State   string
	ZipCode string
}

// Job contains information about the job.
type Job struct {
	Code      string
	Name      string
	Area      string
	Deparment string
}

// Employee contains structured data of a employee.
type Employee struct {
	ID        string
	FirstName string
	LastName  string
	Salary    int64
	Job       *Job
	Address   *Address
}
