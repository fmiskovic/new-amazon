package tests

import (
	"encoding/json"
	"github.com/fmiskovic/new-amz/internal/core/dtos"
	"github.com/fmiskovic/new-amz/internal/core/entities"
	"github.com/fmiskovic/new-amz/internal/core/services"
	"github.com/fmiskovic/new-amz/internal/handlers"
	"github.com/fmiskovic/new-amz/internal/handlers/mappers"
	"github.com/fmiskovic/new-amz/internal/repositories"
	"github.com/fmiskovic/new-amz/internal/validators"
	"github.com/labstack/echo/v4"
	"net/http"
	"net/http/httptest"
	"net/url"
)

func (s *HandlersTestSuite) TestHandleGetItemsPage() {
	e := echo.New()
	e.Validator = validators.New()

	repo := repositories.NewItemRepository(s.testDb.BunDb)
	svc := services.NewItemService(repo)
	handler := handlers.New(
		mappers.NewItemGetPageRequestMapper(),
		mappers.NewItemGetPageResponseMapper(),
		svc.GetPage,
	)

	s.Run("should return items page", func() {
		// given
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		resp := httptest.NewRecorder()
		c := e.NewContext(req, resp)

		// when
		err := handler.Handle(c)

		// then
		s.NoError(err)
		s.Equal(http.StatusOK, resp.Code)
		page := &entities.Page[dtos.ItemDto]{}
		err = json.NewDecoder(resp.Body).Decode(page)
		s.NoError(err)

		s.NotEmpty(page.Elements)
		s.NotEmpty(page.Elements[0].ID)
	})

	s.Run("given page request should return items page", func() {
		// given
		q := make(url.Values)
		q.Set("size", "10")
		q.Set("offset", "0")

		req := httptest.NewRequest(http.MethodGet, "/?"+q.Encode(), nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		resp := httptest.NewRecorder()
		c := e.NewContext(req, resp)

		// when
		err := handler.Handle(c)

		// then
		s.NoError(err)
		s.Equal(http.StatusOK, resp.Code)
		page := &entities.Page[dtos.ItemDto]{}
		err = json.NewDecoder(resp.Body).Decode(page)
		s.NoError(err)

		s.NotEmpty(page.Elements)
		s.NotEmpty(page.Elements[0].ID)
	})
}

func (s *HandlersTestSuite) TestHandleGetItemById() {
	e := echo.New()
	e.Validator = validators.New()

	repo := repositories.NewItemRepository(s.testDb.BunDb)
	svc := services.NewItemService(repo)
	handler := handlers.New(
		mappers.NewItemGetByIdRequestMapper(),
		mappers.NewItemGetByIdResponseMapper(),
		svc.GetById,
	)

	s.Run("should get item by valid id", func() {
		itemId := "200cea28-b2b0-4051-9eb6-9a99e451af01"

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		resp := httptest.NewRecorder()
		c := e.NewContext(req, resp)
		c.SetPath("/item/:id")
		c.SetParamNames("id")
		c.SetParamValues(itemId)

		// when
		err := handler.Handle(c)

		// then
		s.Nil(err)
		s.Equal(http.StatusOK, resp.Code)

		dto := new(dtos.ItemDto)
		err = json.NewDecoder(resp.Body).Decode(dto)
		s.Nil(err)
		s.NotEmpty(dto.ID)
		s.Equal(itemId, dto.ID)
	})

	s.Run("should return 400 when invalid id", func() {
		itemId := "invalid-id"

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		resp := httptest.NewRecorder()
		c := e.NewContext(req, resp)
		c.SetPath("/item/:id")
		c.SetParamNames("id")
		c.SetParamValues(itemId)

		// when
		err := handler.Handle(c)

		// then
		s.NotNil(err)
		s.Equal(http.StatusBadRequest, err.(*echo.HTTPError).Code)
	})
}
