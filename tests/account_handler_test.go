package tests

import (
	"bytes"
	"encoding/json"
	"github.com/fmiskovic/new-amz/internal/core/dtos"
	"github.com/fmiskovic/new-amz/internal/core/entities"
	"github.com/fmiskovic/new-amz/internal/handlers/mappers"
	"net/http"
	"net/http/httptest"
	"net/url"
	"time"

	"github.com/fmiskovic/new-amz/internal/core/services"
	"github.com/fmiskovic/new-amz/internal/handlers"
	"github.com/fmiskovic/new-amz/internal/repositories"
	"github.com/fmiskovic/new-amz/internal/validators"
	"github.com/labstack/echo/v4"
)

func (s *HandlersTestSuite) TestHandleCreateAccount() {
	e := echo.New()
	e.Validator = validators.New()

	repo := repositories.NewAccountRepository(s.testDb.BunDb)
	svc := services.NewAccountService(repo)
	handler := handlers.New(
		mappers.NewCreateAccountRequestMapper(),
		mappers.NewCreateAccountResponseMapper(),
		svc.Create,
	)

	s.Run("should create account", func() {
		// given
		cmd := &dtos.CreateAccountCommand{
			Email:       "fake@mail.com",
			FullName:    "Fake Account",
			DateOfBirth: time.Now(),
			Location:    "Vienna/AUT",
			Gender:      "Male",
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
		s.NoError(err)
		s.Equal(http.StatusCreated, resp.Code)
		answer := new(dtos.CreateAccountAnswer)
		err = json.NewDecoder(resp.Body).Decode(answer)
		s.NoError(err)
		s.NotEmpty(answer.ID)
		s.Equal(cmd.Email, answer.Email)
	})

	s.Run("when email is empty should fail to create account", func() {
		cmd := &dtos.CreateAccountCommand{
			Email:       " ",
			FullName:    "Fake Account",
			DateOfBirth: time.Now(),
			Location:    "Vienna/AUT",
			Gender:      "Male",
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
		s.NotNil(err)
		s.Equal(http.StatusBadRequest, err.(*echo.HTTPError).Code)
	})
}

func (s *HandlersTestSuite) TestHandleGetAccountOrders() {
	e := echo.New()
	e.Validator = validators.New()

	repo := repositories.NewOrderRepository(s.testDb.BunDb)
	svc := services.NewOrderService(repo)
	handler := handlers.New(
		mappers.NewOrderSearchRequestMapper(),
		mappers.NewOrderSearchResponseMapper(),
		svc.Search,
	)

	s.Run("should get account orders", func() {
		// given
		accountId := "220cea28-b2b0-4051-9eb6-9a99e451af01"
		q := make(url.Values)
		q.Set("size", "10")
		q.Set("offset", "0")

		req := httptest.NewRequest(http.MethodGet, "/?"+q.Encode(), nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		resp := httptest.NewRecorder()
		c := e.NewContext(req, resp)
		c.SetPath("/account/:id/orders")
		c.SetParamNames("id")
		c.SetParamValues(accountId)

		// when
		err := handler.Handle(c)

		// then
		s.NoError(err)
		s.Equal(http.StatusOK, resp.Code)
		page := &entities.Page[dtos.OrderDto]{}
		err = json.NewDecoder(resp.Body).Decode(page)
		s.NoError(err)

		s.NotEmpty(page.Elements)
		s.NotEmpty(page.Elements[0].ID)
	})

	s.Run("should fail to get account orders when account id is invalid", func() {
		// given
		accountId := "invalid-id"
		q := make(url.Values)
		q.Set("size", "10")
		q.Set("offset", "0")

		req := httptest.NewRequest(http.MethodGet, "/?"+q.Encode(), nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		resp := httptest.NewRecorder()
		c := e.NewContext(req, resp)
		c.SetPath("/account/:id/orders")
		c.SetParamNames("id")
		c.SetParamValues(accountId)

		// when
		err := handler.Handle(c)

		// then
		s.NotNil(err)
		s.Equal(http.StatusBadRequest, err.(*echo.HTTPError).Code)
	})

	s.Run("should return empty page of account orders when account id is non-existing", func() {
		// given
		accountId := "220cea28-b2b0-4051-9eb6-9a99e451af11"
		q := make(url.Values)
		q.Set("size", "10")
		q.Set("offset", "0")

		req := httptest.NewRequest(http.MethodGet, "/?"+q.Encode(), nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		resp := httptest.NewRecorder()
		c := e.NewContext(req, resp)
		c.SetPath("/account/:id/orders")
		c.SetParamNames("id")
		c.SetParamValues(accountId)

		// when
		err := handler.Handle(c)

		// then
		s.Nil(err)
		s.Equal(http.StatusOK, resp.Code)
		page := &entities.Page[dtos.OrderDto]{}
		err = json.NewDecoder(resp.Body).Decode(page)
		s.NoError(err)
		s.Empty(page.Elements)
		s.Equal(0, page.TotalElements)
	})
}

func (s *HandlersTestSuite) TestHandleGetAccountById() {
	e := echo.New()
	e.Validator = validators.New()

	repo := repositories.NewAccountRepository(s.testDb.BunDb)
	svc := services.NewAccountService(repo)
	handler := handlers.New(
		mappers.NewGetAccountByIdRequestMapper(),
		mappers.NewGetAccountByIdResponseMapper(),
		svc.GetById,
	)

	s.Run("should get account by id", func() {
		// given
		accountId := "220cea28-b2b0-4051-9eb6-9a99e451af01"

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		resp := httptest.NewRecorder()
		c := e.NewContext(req, resp)
		c.SetPath("/account/:id")
		c.SetParamNames("id")
		c.SetParamValues(accountId)

		// when
		err := handler.Handle(c)

		// then
		s.Nil(err)
		s.Equal(http.StatusOK, resp.Code)

		dto := new(dtos.AccountDto)
		err = json.NewDecoder(resp.Body).Decode(dto)
		s.Nil(err)
		s.NotEmpty(dto.ID)
		s.Equal(accountId, dto.ID)
	})

	s.Run("should fail to get account by id when account id is invalid", func() {
		// given
		accountId := "invalid-id"

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		resp := httptest.NewRecorder()
		c := e.NewContext(req, resp)
		c.SetPath("/account/:id")
		c.SetParamNames("id")
		c.SetParamValues(accountId)

		// when
		err := handler.Handle(c)

		// then
		s.NotNil(err)
		s.Equal(http.StatusBadRequest, err.(*echo.HTTPError).Code)
	})

	s.Run("should return 500 when account id is non-existing", func() {
		// given
		accountId := "220cea28-b2b0-4051-9eb6-9a99e451af11"
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		resp := httptest.NewRecorder()
		c := e.NewContext(req, resp)
		c.SetPath("/account/:id")
		c.SetParamNames("id")
		c.SetParamValues(accountId)

		// when
		err := handler.Handle(c)

		// then
		s.NotNil(err)
		s.Equal(http.StatusInternalServerError, err.(*echo.HTTPError).Code)
	})

}
