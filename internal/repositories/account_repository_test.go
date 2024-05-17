package repositories

import (
	"github.com/fmiskovic/new-amz/internal/core/entities"
	"github.com/google/uuid"
)

func (s *RepositoryTestSuite) TestGetAccountById() {
	repo := NewAccountRepository(s.testDb.BunDb)

	s.Run("should return account by id", func() {
		// given
		accountId := uuid.MustParse("220cea28-b2b0-4051-9eb6-9a99e451af01")
		// when
		acc, err := repo.GetById(s.testDb.Ctx, accountId)
		// then
		s.Nil(err)
		s.Equal(accountId, acc.ID)
		s.Equal("john@smith.com", acc.Email)
	})

	s.Run("should return error if account not found", func() {
		// given
		accountId := uuid.MustParse("220cea28-b2b0-4051-9eb6-9a99e451af10")
		// when
		_, err := repo.GetById(s.testDb.Ctx, accountId)
		// assert
		s.NotNil(err)
	})
}

func (s *RepositoryTestSuite) TestGetAccountByEmail() {
	repo := NewAccountRepository(s.testDb.BunDb)

	s.Run("should return account by email", func() {
		// given
		email := "john@smith.com"
		// when
		acc, err := repo.GetByEmail(s.testDb.Ctx, email)
		// then
		s.Nil(err)
		s.Equal(email, acc.Email)
	})

	s.Run("should return error if account not found", func() {
		// given
		email := "fake@mail.com"
		// when
		_, err := repo.GetByEmail(s.testDb.Ctx, email)
		// assert
		s.NotNil(err)
	})
}

func (s *RepositoryTestSuite) TestCreateAccount() {
	repo := NewAccountRepository(s.testDb.BunDb)

	s.Run("should create new account", func() {
		// given
		acc := entities.NewAccountBuilder().
			Email("new1@mail.com").
			Build()
		// when
		err := repo.Create(s.testDb.Ctx, acc)
		// then
		s.Nil(err)
		s.NotNil(acc.ID)
		s.NotNil(acc.CreatedAt)
	})

	s.Run("should return error if account already exists", func() {
		// given
		acc := entities.NewAccountBuilder().
			Email("john@smith.com").
			Build()
		// when
		err := repo.Create(s.testDb.Ctx, acc)
		// assert
		s.NotNil(err)
	})
}
