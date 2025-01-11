package commands

import (
	"go-clean-architecture-sample/application/common/interfaces"
)

type DeleteEmployeeByIdCommand struct {
	EmployeeId int
}

type DeleteEmployeeByIdCommandHandler struct {
	employeeRepository interfaces.EmployeeRepository
}

func NewDeleteEmployeeByIdCommandHandler(
	employeeRepository interfaces.EmployeeRepository,
) *DeleteEmployeeByIdCommandHandler {
	return &DeleteEmployeeByIdCommandHandler{
		employeeRepository: employeeRepository,
	}
}

func (h *DeleteEmployeeByIdCommandHandler) Handle(request DeleteEmployeeByIdCommand) (err error) {
	if err = h.employeeRepository.DeleteById(request.EmployeeId); err != nil {
		return err
	}

	return nil
}
