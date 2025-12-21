package company

import (
	"advance/company/employee"
	"advance/company/employee/employees"
	"fmt"
)

func GetCompany() *Company {
	return &Company{make(map[uint64]employee.Employee), 0}
}

type Company struct {
	employees map[uint64]employee.Employee
	counter   uint64
}

func (c *Company) AddEmployee(e employee.Employee) {
	c.employees[c.counter] = e
	c.counter++
}

func (c *Company) ListEmployees() {
	if len(c.employees) == 0 {
		fmt.Println("No Employees found")
		return
	}
	for key, value := range c.employees {
		fmt.Printf("ID %d: ", key)
		value.GetDetails()
	}
}

func Start() {
	company := GetCompany()

	company.ListEmployees()
	company.AddEmployee(&employees.FullTime{})
	company.AddEmployee(&employees.PartTime{})
	company.ListEmployees()
}
