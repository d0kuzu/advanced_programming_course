package company

import (
	employees2 "advance/assignment1/company/employee/employees"
)

func Start() {
	company := GetCompany()

	company.ListEmployees()
	company.AddEmployee(&employees2.FullTime{})
	company.AddEmployee(&employees2.PartTime{})
	company.ListEmployees()
}
