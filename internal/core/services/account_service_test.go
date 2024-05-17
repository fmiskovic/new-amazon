package services

import (
	"context"
	"github.com/fmiskovic/new-amz/internal/core/dtos"
	"github.com/fmiskovic/new-amz/internal/core/entities"
	"github.com/fmiskovic/new-amz/internal/core/repositories"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

func TestCreateAccount(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	t.Run("create new account with valid data should return account dto", func(t *testing.T) {
		repoMock := repositories.NewAccountRepositoryMock[uuid.UUID](t)
		svc := NewAccountService(repoMock)

		cmd := dtos.CreateAccountCommand{
			Email:       "fake@mail.com",
			FullName:    "Fake Name",
			Location:    "Fake Location",
			DateOfBirth: time.Date(1990, time.January, 1, 0, 0, 0, 0, time.UTC),
			Gender:      "Male",
		}

		repoMock.On("Create", mock.Anything, mock.Anything).Return(nil).Once()

		got, err := svc.Create(ctx, cmd)
		assert.NoError(t, err)
		assert.Equal(t, cmd.Email, got.Email)
		assert.Equal(t, cmd.Gender, got.Gender)
		repoMock.AssertCalled(t, "Create", mock.Anything, mock.Anything)
	})

	t.Run("create new account with empty email should return error", func(t *testing.T) {
		repoMock := repositories.NewAccountRepositoryMock[uuid.UUID](t)
		svc := NewAccountService(repoMock)

		cmd := dtos.CreateAccountCommand{
			Email:       " ",
			FullName:    "Fake Name",
			Location:    "Fake Location",
			DateOfBirth: time.Date(1990, time.January, 1, 0, 0, 0, 0, time.UTC),
			Gender:      "Male",
		}

		_, err := svc.Create(ctx, cmd)
		assert.NotNil(t, err)
		repoMock.AssertNotCalled(t, "Create", mock.Anything, mock.Anything)
	})

	t.Run("create new account with existing email should return error", func(t *testing.T) {
		repoMock := repositories.NewAccountRepositoryMock[uuid.UUID](t)
		svc := NewAccountService(repoMock)

		cmd := dtos.CreateAccountCommand{
			Email:       "fake@mail.com",
			FullName:    "Fake Name",
			Location:    "Fake Location",
			DateOfBirth: time.Date(1990, time.January, 1, 0, 0, 0, 0, time.UTC),
			Gender:      "Male",
		}

		repoMock.On("Create", mock.Anything, mock.Anything).Return(ErrorEmailNotUnique).Once()

		_, err := svc.Create(ctx, cmd)
		assert.NotNil(t, err)
		repoMock.AssertCalled(t, "Create", mock.Anything, mock.Anything)
	})
}

func TestGetAccountById(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	t.Run("get account by id should return account dto", func(t *testing.T) {
		repoMock := repositories.NewAccountRepositoryMock[uuid.UUID](t)
		svc := NewAccountService(repoMock)

		a := entities.NewAccountBuilder().
			Email("fake@mail.com").
			FullName("Fake Name").
			Location("Fake Location").
			DateOfBirth(time.Date(1990, time.January, 1, 0, 0, 0, 0, time.UTC)).
			Gender(entities.MALE).
			Build()

		repoMock.On("GetById", mock.Anything, mock.Anything).Return(*a, nil).Once()

		got, err := svc.GetById(ctx, uuid.New())
		assert.NoError(t, err)
		assert.Equal(t, a.Email, got.Email)
		assert.Equal(t, dtos.GenderDto(a.Gender.Stringify()), got.Gender)
		repoMock.AssertCalled(t, "GetById", mock.Anything, mock.Anything)
	})

	t.Run("get account by non existing id should return error", func(t *testing.T) {
		repoMock := repositories.NewAccountRepositoryMock[uuid.UUID](t)
		svc := NewAccountService(repoMock)

		repoMock.On("GetById", mock.Anything, mock.Anything).Return(entities.Account{}, entities.ErrorEntityNotFound).Once()

		_, err := svc.GetById(ctx, uuid.New())
		assert.NotNil(t, err)
		repoMock.AssertCalled(t, "GetById", mock.Anything, mock.Anything)
	})
}
