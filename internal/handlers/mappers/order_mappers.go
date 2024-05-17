package mappers

import (
	"github.com/fmiskovic/new-amz/internal/core/dtos"
	"github.com/fmiskovic/new-amz/internal/core/entities"
	"github.com/fmiskovic/new-amz/internal/handlers"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type OrderGetByIdRequestMapper struct{}

func NewOrderGetByIdRequestMapper() OrderGetByIdRequestMapper {
	return OrderGetByIdRequestMapper{}
}

func (m OrderGetByIdRequestMapper) Map(c echo.Context) (uuid.UUID, error) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return id, handlers.NewErr("failed to parse order id", err, 400)
	}
	return id, nil
}

type OrderGetByIdResponseMapper struct{}

func NewOrderGetByIdResponseMapper() OrderGetByIdResponseMapper {
	return OrderGetByIdResponseMapper{}
}

func (m OrderGetByIdResponseMapper) Map(c echo.Context, out dtos.OrderDto) error {
	return c.JSON(200, out)
}

type OrderSearchRequestMapper struct{}

func NewOrderSearchRequestMapper() OrderSearchRequestMapper {
	return OrderSearchRequestMapper{}
}

func (m OrderSearchRequestMapper) Map(c echo.Context) (dtos.OrderFilter, error) {
	var filter dtos.OrderFilter

	accountId, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return filter, handlers.NewErr("failed to parse account id", err, 400)
	}

	filter.AccountID = accountId
	filter.PageRequest = pageRequestMapper(c)

	return filter, nil
}

type OrderSearchResponseMapper struct{}

func NewOrderSearchResponseMapper() OrderSearchResponseMapper {
	return OrderSearchResponseMapper{}
}

func (m OrderSearchResponseMapper) Map(c echo.Context, out entities.Page[dtos.OrderDto]) error {
	return c.JSON(200, out)
}

type OrderCreateRequestMapper struct{}

func NewOrderCreateRequestMapper() OrderCreateRequestMapper {
	return OrderCreateRequestMapper{}
}

func (m OrderCreateRequestMapper) Map(c echo.Context) (dtos.CreateOrderCommand, error) {
	var cmd dtos.CreateOrderCommand
	if err := c.Bind(&cmd); err != nil {
		return cmd, handlers.NewErr("failed to bind create order command request", err, 400)
	}
	return cmd, nil
}

type OrderCreateResponseMapper struct{}

func NewOrderCreateResponseMapper() OrderCreateResponseMapper {
	return OrderCreateResponseMapper{}
}

func (m OrderCreateResponseMapper) Map(c echo.Context, out dtos.CreateOrderAnswer) error {
	return c.JSON(201, out)
}
