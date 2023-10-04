package service

import (
	"errors"
	"restEcho1/features/users"
	"restEcho1/features/users/mocks"
	helper "restEcho1/helper/mocks"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {
	generator := helper.NewGeneratorInterface(t)
	data := mocks.NewUserDataInterface(t)
	service := New(data, generator)
	newUser := users.User{
		Nama:     "jerry",
		HP:       "12345",
		Password: "mantul123",
	}

	t.Run("Success insert", func(t *testing.T) {
		generator.On("GenerateUUID").Return("randomUUID", nil).Once()
		newUser.ID = "randomUUID"
		data.On("Insert", newUser).Return(newUser, nil).Once()

		result, err := service.Register(newUser)
		assert.Nil(t, err)
		assert.Equal(t, newUser.ID, result.ID)
		assert.Equal(t, newUser.Nama, result.Nama)
		generator.AssertExpectations(t)
		data.AssertExpectations(t)
	})

	t.Run("Generate failed", func(t *testing.T) {
		generator.On("GenerateUUID").Return("", errors.New("some error on generator")).Once()

		result, err := service.Register(newUser)
		assert.Error(t, err)
		assert.EqualError(t, err, "id generator failed")
		assert.Equal(t, users.User{}, result)
		generator.AssertExpectations(t)
	})

	// t.Run("Insert failed", func(t *testing.T) {

	// })
}
