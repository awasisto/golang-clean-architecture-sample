package interfaces

import "go-clean-architecture-sample/domain/entities"

type EmployeeRepository interface {
	Add(employee entities.Employee) (employeeId int, err error)
	GetAll() (employees []entities.Employee, err error)
	GetById(id int) (employee *entities.Employee, err error)
	DeleteById(id int) (err error)
}
