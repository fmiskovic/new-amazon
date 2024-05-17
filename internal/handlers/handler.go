package handlers

import (
	"errors"
	"net/http"

	"github.com/fmiskovic/new-amz/internal/core"
	"github.com/labstack/echo/v4"
)

// RequestMapper is interfaces that define how to map between HTTP requests and the core request.
type RequestMpper[T any] interface {
	Map(c echo.Context) (T, error)
}

// ResponseMapper is interfaces that define how to map between core responses and HTTP responses.
type ResponseMapper[T any] interface {
	Map(c echo.Context, t T) error
}

// Handler is a struct that contains the service function, request mapper, and response mapper.
type Handler[In any, Out any] struct {
	serviceFunc    core.ServiceFunc[In, Out]
	requestMapper  RequestMpper[In]
	responseMapper ResponseMapper[Out]
}

// New creates a new handler.
// It takes a request mapper, response mapper, and service function.
// Service function is a target function that takes a context and a request and returns a response and an error.
func New[In any, Out any](
	reqMapper RequestMpper[In],
	resMapper ResponseMapper[Out],
	svcFunc core.ServiceFunc[In, Out],
) Handler[In, Out] {
	return Handler[In, Out]{
		serviceFunc:    svcFunc,
		requestMapper:  reqMapper,
		responseMapper: resMapper,
	}
}

// Handle is the generic receiver that handles incoming HTTP request.
func (h Handler[In, Out]) Handle(c echo.Context) error {
	// Parse request
	in, err := h.requestMapper.Map(c)
	if err != nil {
		c.Logger().Error(err)
		var herr HandlerError
		if errors.As(err, &herr) {
			return echo.NewHTTPError(herr.code, herr.Error())
		}
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// Validate request
	err = c.Validate(in)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// Call out to service function
	out, err := h.serviceFunc(c.Request().Context(), in)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	// Map and return response
	return h.responseMapper.Map(c, out)
}
