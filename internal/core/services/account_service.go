package services

import (
	"context"
	"errors"
	"fmt"
	"github.com/fmiskovic/new-amz/internal/core/dtos"
	"github.com/fmiskovic/new-amz/internal/core/entities"
	"github.com/fmiskovic/new-amz/internal/core/repositories"
	"github.com/google/uuid"
	"strings"
)

var (
	ErrorEmailRequired  = errors.New("email is required")
	ErrorEmailNotUnique = errors.New("email is not unique")
)

// AccountService represents business logic related to entities.Account.
type AccountService struct {
	repo repositories.AccountRepository[uuid.UUID]
}

// NewAccountService instantiates new Account Service.
func NewAccountService(repo repositories.AccountRepository[uuid.UUID]) AccountService {
	return AccountService{repo}
}

// Create creates new account.
func (s AccountService) Create(ctx context.Context, cmd dtos.CreateAccountCommand) (dtos.CreateAccountAnswer, error) {
	if len(strings.TrimSpace(cmd.Email)) == 0 {
		return dtos.CreateAccountAnswer{}, ErrorEmailRequired
	}
	a := entities.NewAccountBuilder().
		Email(cmd.Email).
		FullName(cmd.FullName).
		DateOfBirth(cmd.DateOfBirth).
		Location(cmd.Location).
		Gender(cmd.Gender.Numberfy()).
		Build()

	if err := s.repo.Create(ctx, a); err != nil {
		return dtos.CreateAccountAnswer{}, newError("failed to create account", err)
	}

	return dtos.CreateAccountAnswer{AccountDto: dtos.ToAccountDto(*a)}, nil
}

// GetById returns existing account.
func (s AccountService) GetById(ctx context.Context, id uuid.UUID) (dtos.AccountDto, error) {
	a, err := s.repo.GetById(ctx, id)
	if err != nil {
		return dtos.AccountDto{}, newError(fmt.Sprintf("failed to get account by id: %s", id.String()), err)
	}
	return dtos.ToAccountDto(a), nil
}
