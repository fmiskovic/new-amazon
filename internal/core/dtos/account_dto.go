package dtos

import (
	"github.com/fmiskovic/new-amz/internal/core/entities"
	"time"
)

type AccountDto struct {
	ID          string    `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Email       string    `validate:"required,min=3" json:"email"`
	FullName    string    `json:"full_name"`
	DateOfBirth time.Time `json:"date_of_birth"`
	Location    string    `json:"location"`
	Gender      GenderDto `json:"gender"`
}

// GenderDto can be Male, Female and Other.
type GenderDto string

// Numberfy converts GenderDto into a account.Gender.
func (g GenderDto) Numberfy() entities.Gender {
	switch g {
	case "Male":
		return entities.MALE
	case "Female":
		return entities.FEMALE
	case "Other":
		return entities.OTHER
	default:
		return entities.OTHER
	}
}

// ToAccountDto converts Account entity into a Account DTO.
func ToAccountDto(a entities.Account) AccountDto {
	return AccountDto{
		ID:          a.ID.String(),
		Email:       a.Email,
		FullName:    a.FullName,
		DateOfBirth: a.DateOfBirth,
		Location:    a.Location,
		Gender:      GenderDto(a.Gender.Stringify()),
	}
}

type CreateAccountCommand struct {
	Email       string    `validate:"required,min=3" json:"email"`
	FullName    string    `json:"full_name"`
	DateOfBirth time.Time `json:"date_of_birth"`
	Location    string    `json:"location"`
	Gender      GenderDto `json:"gender"`
}

// CreateAccountAnswer is a response to a create Command
type CreateAccountAnswer struct {
	AccountDto
}
