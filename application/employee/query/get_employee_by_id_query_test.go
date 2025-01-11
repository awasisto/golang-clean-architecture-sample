package query

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go-clean-architecture-sample/application/common/interfaces/mocks"
	"go-clean-architecture-sample/domain/entities"
	"testing"
)

func TestGetEmployeeById(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	mockEmployeeRepository := mocks.NewMockEmployeeRepository(ctrl)
	mockAvatarProvider := mocks.NewMockAvatarProvider(ctrl)

	getEmployeeByIdQueryHandler := NewGetEmployeeByIdQueryHandler(
		mockEmployeeRepository,
		mockAvatarProvider,
	)

	mockEmployeeRepository.EXPECT().
		GetById(42).
		Return(
			&entities.Employee{
				Id:    42,
				Name:  "John Smith",
				Email: "john.smith@example.com",
			},
			nil,
		)

	mockAvatarProvider.EXPECT().
		GetAvatarUrlByEmail("john.smith@example.com").
		Return("http://example.com/john_smith.jpg", nil)

	want := entities.Employee{
		Id:        42,
		Name:      "John Smith",
		Email:     "john.smith@example.com",
		AvatarUrl: "http://example.com/john_smith.jpg",
	}

	got, _ := getEmployeeByIdQueryHandler.Handle(GetEmployeeByIdQuery{EmployeeId: 42})

	assert.Equal(t, want, *got)
}
