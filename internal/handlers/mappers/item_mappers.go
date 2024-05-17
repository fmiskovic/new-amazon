package mappers

import (
	"github.com/fmiskovic/new-amz/internal/core/dtos"
	"github.com/fmiskovic/new-amz/internal/core/entities"
	"github.com/fmiskovic/new-amz/internal/handlers"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type ItemGetByIdRequestMapper struct{}

func NewItemGetByIdRequestMapper() ItemGetByIdRequestMapper {
	return ItemGetByIdRequestMapper{}
}

func (m ItemGetByIdRequestMapper) Map(c echo.Context) (uuid.UUID, error) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return id, handlers.NewErr("failed to parse item id", err, 400)
	}
	return id, nil
}

type ItemGetByIdResponseMapper struct{}

func NewItemGetByIdResponseMapper() ItemGetByIdResponseMapper {
	return ItemGetByIdResponseMapper{}
}

func (m ItemGetByIdResponseMapper) Map(c echo.Context, out dtos.ItemDto) error {
	return c.JSON(200, out)
}

type ItemGetPageRequestMapper struct{}

func NewItemGetPageRequestMapper() ItemGetPageRequestMapper {
	return ItemGetPageRequestMapper{}
}

func (m ItemGetPageRequestMapper) Map(c echo.Context) (entities.Pageable, error) {
	return pageRequestMapper(c), nil
}

type ItemGetPageResponseMapper struct{}

func NewItemGetPageResponseMapper() ItemGetPageResponseMapper {
	return ItemGetPageResponseMapper{}
}

func (m ItemGetPageResponseMapper) Map(c echo.Context, out entities.Page[dtos.ItemDto]) error {
	return c.JSON(200, out)
}
