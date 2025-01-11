package database

import (
	"go-clean-architecture-sample/application/common/errors"
	"go-clean-architecture-sample/domain/entities"
)

type InMemoryEmployeeDatabase struct {
	employees map[int]entities.Employee
	nextId    int
}

func NewInMemoryEmployeeDatabase() *InMemoryEmployeeDatabase {
	return &InMemoryEmployeeDatabase{
		employees: make(map[int]entities.Employee),
		nextId:    1,
	}
}

func (d *InMemoryEmployeeDatabase) Add(employee entities.Employee) (int, error) {
	employee.Id = d.nextId
	d.employees[d.nextId] = employee
	d.nextId++
	return employee.Id, nil
}

func (d *InMemoryEmployeeDatabase) GetAll() ([]entities.Employee, error) {
	employees := make([]entities.Employee, 0, len(d.employees))
	for _, employee := range d.employees {
		employees = append(employees, employee)
	}
	return employees, nil
}

func (d *InMemoryEmployeeDatabase) GetById(id int) (*entities.Employee, error) {
	employee, exists := d.employees[id]
	if !exists {
		return nil, errors.ErrNotFound
	}
	return &employee, nil
}

func (d *InMemoryEmployeeDatabase) DeleteById(id int) error {
	_, exists := d.employees[id]
	if !exists {
		return errors.ErrNotFound
	}
	delete(d.employees, id)
	return nil
}
