package company

import "advance/company/employee/employees"

func Start() {
	company := GetCompany()

	company.ListEmployees()
	company.AddEmployee(&employees.FullTime{})
	company.AddEmployee(&employees.PartTime{})
	company.ListEmployees()
}
