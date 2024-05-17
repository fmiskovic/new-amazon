package entities

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
	"time"
)

// Account will store information about each user account.
type Account struct {
	bun.BaseModel `bun:"table:accounts,alias:a"`

	Entity
	Email       string    `bun:"email,notnull,unique"`
	FullName    string    `bun:"full_name,nullzero"`
	DateOfBirth time.Time `bun:"date_of_birth,nullzero"`
	Location    string    `bun:"location,nullzero"`
	Gender      Gender    `bun:"gender,nullzero"`

	// one-to-many relation
	Orders []*Order `bun:"rel:has-many,join:id=account_id"`
}

type AccountBuilder struct {
	email       string
	fullName    string
	dateOfBirth time.Time
	location    string
	gender      Gender
}

func NewAccountBuilder() *AccountBuilder {
	return &AccountBuilder{}
}

// Email sets the email on the Builder.
func (b *AccountBuilder) Email(email string) *AccountBuilder {
	b.email = email
	return b
}

// FullName sets the full name on the Builder.
func (b *AccountBuilder) FullName(fullName string) *AccountBuilder {
	b.fullName = fullName
	return b
}

// DateOfBirth sets the date of birth on the Builder.
func (b *AccountBuilder) DateOfBirth(dateOfBirth time.Time) *AccountBuilder {
	b.dateOfBirth = dateOfBirth
	return b
}

// Location sets the location on the Builder.
func (b *AccountBuilder) Location(location string) *AccountBuilder {
	b.location = location
	return b
}

// Gender sets the gender on the Builder.
func (b *AccountBuilder) Gender(gender Gender) *AccountBuilder {
	b.gender = gender
	return b
}

// Build constructs an Account instance from the Builder.
func (b *AccountBuilder) Build() *Account {
	return &Account{
		Entity:      Entity{ID: uuid.New(), CreatedAt: time.Now(), UpdatedAt: time.Now()},
		Email:       b.email,
		FullName:    b.fullName,
		DateOfBirth: b.dateOfBirth,
		Location:    b.location,
		Gender:      b.gender,
	}
}

// Gender is either MALE, FEMALE or OTHER.
type Gender uint8

// Stringify converts Gender into string.
func (g Gender) Stringify() string {
	switch g {
	case 0:
		return "Male"
	case 1:
		return "Female"
	case 2:
		return "Other"
	default:
		return "Other"
	}
}

const (
	MALE Gender = iota
	FEMALE
	OTHER
)
