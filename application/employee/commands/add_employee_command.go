package commands

import (
	"go-clean-architecture-sample/application/common/interfaces"
	"go-clean-architecture-sample/domain/entities"
)

type AddEmployeeCommand struct {
	Name  string
	Email string
}

type AddEmployeeCommandHandler struct {
	employeeRepository interfaces.EmployeeRepository
	avatarProvider     interfaces.AvatarProvider
}

func NewAddEmployeeCommandHandler(
	employeeRepository interfaces.EmployeeRepository,
	avatarProvider interfaces.AvatarProvider,
) *AddEmployeeCommandHandler {
	return &AddEmployeeCommandHandler{
		employeeRepository: employeeRepository,
		avatarProvider:     avatarProvider,
	}
}

func (h *AddEmployeeCommandHandler) Handle(request AddEmployeeCommand) (createdEmployee *entities.Employee, err error) {
	entity := entities.Employee{
		Name:  request.Name,
		Email: request.Email,
	}

	entity.Id, err = h.employeeRepository.Add(entity)
	if err != nil {
		return nil, err
	}

	entity.AvatarUrl, err = h.avatarProvider.GetAvatarUrlByEmail(entity.Email)
	if err != nil {
		return nil, err
	}

	return &entity, nil
}
