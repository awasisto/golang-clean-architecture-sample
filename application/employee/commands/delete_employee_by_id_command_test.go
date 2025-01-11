package commands

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go-clean-architecture-sample/application/common/interfaces/mocks"
	"testing"
)

func TestDeleteEmployeeById(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	mockEmployeeRepository := mocks.NewMockEmployeeRepository(ctrl)

	deleteEmployeeByIdCommandHandler := NewDeleteEmployeeByIdCommandHandler(
		mockEmployeeRepository,
	)

	mockEmployeeRepository.EXPECT().
		DeleteById(42).
		Return(nil)

	got := deleteEmployeeByIdCommandHandler.Handle(DeleteEmployeeByIdCommand{
		EmployeeId: 42,
	})

	assert.Equal(t, nil, got)
}
