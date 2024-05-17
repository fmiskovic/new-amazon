package tests

import (
	"bytes"
	"encoding/json"
	"github.com/fmiskovic/new-amz/internal/core/dtos"
	"github.com/fmiskovic/new-amz/internal/core/services"
	"github.com/fmiskovic/new-amz/internal/handlers"
	"github.com/fmiskovic/new-amz/internal/handlers/mappers"
	"github.com/fmiskovic/new-amz/internal/repositories"
	"github.com/fmiskovic/new-amz/internal/validators"
	"github.com/labstack/echo/v4"
	"net/http"
	"net/http/httptest"
)

func (s *HandlersTestSuite) TestHandleCreateOrder() {
	e := echo.New()
	e.Validator = validators.New()

	repo := repositories.NewOrderRepository(s.testDb.BunDb)
	svc := services.NewOrderService(repo)
	handler := handlers.New(
		mappers.NewOrderCreateRequestMapper(),
		mappers.NewOrderCreateResponseMapper(),
		svc.Create,
	)

	s.Run("should create order", func() {
		// given
		cmd := &dtos.CreateOrderCommand{
			AccountID: "220cea28-b2b0-4051-9eb6-9a99e451af01",
			Items: []dtos.OrderItemDto{
				{
					ItemID:   "200cea28-b2b0-4051-9eb6-9a99e451af01",
					Quantity: 2,
				},
			},
		}
		b, err := json.Marshal(cmd)
		if err != nil {
			s.Error(err)
		}
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(b))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		resp := httptest.NewRecorder()
		c := e.NewContext(req, resp)

		// when
		err = handler.Handle(c)

		// then
		s.Nil(err)
		s.Equal(http.StatusCreated, resp.Code)
		answer := new(dtos.CreateOrderAnswer)
		err = json.NewDecoder(resp.Body).Decode(answer)
		s.Nil(err)
		s.NotEmpty(answer.ID)
		s.Equal(cmd.AccountID, answer.AccountID)
		s.NotEmpty(answer.Items)
	})

	s.Run("should return 500 when non-existing account id", func() {
		// given
		cmd := &dtos.CreateOrderCommand{
			AccountID: "220cea28-b2b0-4051-9eb6-9a99e451af11",
			Items: []dtos.OrderItemDto{
				{
					ItemID:   "200cea28-b2b0-4051-9eb6-9a99e451af01",
					Quantity: 2,
				},
			},
		}
		b, err := json.Marshal(cmd)
		if err != nil {
			s.Error(err)
		}
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(b))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		resp := httptest.NewRecorder()
		c := e.NewContext(req, resp)

		// when
		err = handler.Handle(c)
		s.NotNil(err)
		s.Equal(http.StatusInternalServerError, err.(*echo.HTTPError).Code)
	})

	s.Run("should return 500 when non-existing item id", func() {
		// given
		cmd := &dtos.CreateOrderCommand{
			AccountID: "220cea28-b2b0-4051-9eb6-9a99e451af01",
			Items: []dtos.OrderItemDto{
				{
					ItemID:   "200cea28-b2b0-4051-9eb6-9a99e451af11",
					Quantity: 2,
				},
			},
		}
		b, err := json.Marshal(cmd)
		if err != nil {
			s.Error(err)
		}
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(b))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		resp := httptest.NewRecorder()
		c := e.NewContext(req, resp)

		// when
		err = handler.Handle(c)
		s.NotNil(err)
		s.Equal(http.StatusInternalServerError, err.(*echo.HTTPError).Code)
	})
}

func (s *HandlersTestSuite) TestHandleGetOrderById() {
	e := echo.New()
	e.Validator = validators.New()

	repo := repositories.NewOrderRepository(s.testDb.BunDb)
	svc := services.NewOrderService(repo)
	handler := handlers.New(
		mappers.NewOrderGetByIdRequestMapper(),
		mappers.NewOrderGetByIdResponseMapper(),
		svc.GetById,
	)

	s.Run("should get order by valid id", func() {
		orderId := "210cea28-b2b0-4051-9eb6-9a99e451af01"

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		resp := httptest.NewRecorder()
		c := e.NewContext(req, resp)
		c.SetPath("/order/:id")
		c.SetParamNames("id")
		c.SetParamValues(orderId)

		// when
		err := handler.Handle(c)

		// then
		s.Nil(err)
		s.Equal(http.StatusOK, resp.Code)

		dto := new(dtos.OrderDto)
		err = json.NewDecoder(resp.Body).Decode(dto)
		s.Nil(err)
		s.NotEmpty(dto.ID)
		s.Equal(orderId, dto.ID)
		s.NotEmpty(dto.Items)
	})

	s.Run("should return 400 when invalid id", func() {
		orderId := "invalid-id"

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		resp := httptest.NewRecorder()
		c := e.NewContext(req, resp)
		c.SetPath("/order/:id")
		c.SetParamNames("id")
		c.SetParamValues(orderId)

		// when
		err := handler.Handle(c)

		// then
		s.NotNil(err)
		s.Equal(http.StatusBadRequest, err.(*echo.HTTPError).Code)
	})

	s.Run("should return 500 when non-existing id", func() {
		orderId := "210cea28-b2b0-4051-9eb6-9a99e451af11"

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		resp := httptest.NewRecorder()
		c := e.NewContext(req, resp)
		c.SetPath("/order/:id")
		c.SetParamNames("id")
		c.SetParamValues(orderId)

		// when
		err := handler.Handle(c)

		// then
		s.NotNil(err)
		s.Equal(http.StatusInternalServerError, err.(*echo.HTTPError).Code)
	})
}
