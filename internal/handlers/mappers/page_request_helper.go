package mappers

import (
	"github.com/fmiskovic/new-amz/internal/core/entities"
	"github.com/labstack/echo/v4"
	"strconv"
	"strings"
)

func pageRequestMapper(c echo.Context) entities.Pageable {
	size, err := strconv.Atoi(c.QueryParam("size"))
	if err != nil {
		size = 10
	}
	offset, err := strconv.Atoi(c.QueryParam("offset"))
	if err != nil {
		offset = 0
	}

	pageReq := entities.Pageable{
		Size:   size,
		Offset: offset,
		Sort:   resolveSort(c),
	}
	return pageReq
}

func resolveSort(c echo.Context) entities.Sort {
	sortParam := c.QueryParam("sort")
	if sortParam == "" {
		return entities.Sort{}
	}

	// split the sort parameter into individual sort orderParams
	orderParams := strings.Split(sortParam, ",")

	var sortOrders []*entities.SortOrder

	for i := range orderParams {
		o := strings.Split(strings.TrimSpace(orderParams[i]), " ")
		sortOrder := entities.NewSortOrder(entities.WithProperty(o[0]), entities.WithDirection(entities.ASC))
		if len(o) == 2 {
			sortOrder.Direction = entities.Direction(o[1])
		}
		sortOrders = append(sortOrders, sortOrder)
	}

	return entities.NewSort(sortOrders...)
}
