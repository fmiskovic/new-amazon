package mappers

import (
	"github.com/fmiskovic/new-amz/internal/core/dtos"
	"github.com/fmiskovic/new-amz/internal/handlers"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type CreateAccountRequestMapper struct{}

func NewCreateAccountRequestMapper() CreateAccountRequestMapper {
	return CreateAccountRequestMapper{}
}

func (m CreateAccountRequestMapper) Map(c echo.Context) (dtos.CreateAccountCommand, error) {
	cmd := new(dtos.CreateAccountCommand)
	if err := c.Bind(cmd); err != nil {
		return *cmd, handlers.NewErr("failed to bind create account request", err, 400)
	}
	return *cmd, nil
}

type CreateAccountResponseMapper struct{}

func NewCreateAccountResponseMapper() CreateAccountResponseMapper {
	return CreateAccountResponseMapper{}
}

func (m CreateAccountResponseMapper) Map(c echo.Context, out dtos.CreateAccountAnswer) error {
	return c.JSON(201, out)
}

type GetAccountByIdRequestMapper struct{}

func NewGetAccountByIdRequestMapper() GetAccountByIdRequestMapper {
	return GetAccountByIdRequestMapper{}
}

func (m GetAccountByIdRequestMapper) Map(c echo.Context) (uuid.UUID, error) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return id, handlers.NewErr("failed to parse account id", err, 400)
	}
	return id, nil
}

type GetAccountByIdResponseMapper struct{}

func NewGetAccountByIdResponseMapper() GetAccountByIdResponseMapper {
	return GetAccountByIdResponseMapper{}
}

func (m GetAccountByIdResponseMapper) Map(c echo.Context, out dtos.AccountDto) error {
	return c.JSON(200, out)
}
