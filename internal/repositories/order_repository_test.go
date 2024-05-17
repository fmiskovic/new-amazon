package repositories

import (
	"github.com/fmiskovic/new-amz/internal/core/entities"
	"github.com/google/uuid"
)

func (s *RepositoryTestSuite) TestGetOrderById() {
	repo := NewOrderRepository(s.testDb.BunDb)

	s.Run("should return order by id", func() {
		// given
		orderId := uuid.MustParse("210cea28-b2b0-4051-9eb6-9a99e451af01")
		// when
		order, err := repo.GetById(s.testDb.Ctx, orderId)
		// then
		s.Nil(err)
		s.Equal("220cea28-b2b0-4051-9eb6-9a99e451af01", order.AccountID.String())
	})

	s.Run("should return error if order not found", func() {
		// given
		orderId := uuid.MustParse("210cea28-b2b0-4051-9eb6-9a99e451af10")
		// when
		_, err := repo.GetById(s.testDb.Ctx, orderId)
		// then
		s.NotNil(err)
	})
}

func (s *RepositoryTestSuite) TestSearchAccountOrders() {
	repo := NewOrderRepository(s.testDb.BunDb)

	s.Run("should return page of account orders", func() {
		// given
		accountId := uuid.MustParse("220cea28-b2b0-4051-9eb6-9a99e451af01")
		pageRequest := entities.Pageable{
			Size:   2,
			Offset: 0,
		}
		// when
		page, err := repo.Search(s.testDb.Ctx, accountId, pageRequest)
		// then
		s.Nil(err)
		s.NotEmpty(page.Elements)
		s.Equal(accountId, page.Elements[0].AccountID)
	})

	s.Run("given zero page request should return page of account orders", func() {
		// given
		accountId := uuid.MustParse("220cea28-b2b0-4051-9eb6-9a99e451af01")
		pageRequest := entities.Pageable{}
		// when
		page, err := repo.Search(s.testDb.Ctx, accountId, pageRequest)
		// then
		s.Nil(err)
		s.NotEmpty(page.Elements)
		s.Equal(accountId, page.Elements[0].AccountID)
	})

	s.Run("given invalid account id should return empty page", func() {
		// given
		accountId := uuid.MustParse("220cea28-b2b0-4051-9eb6-9a99e451af10")
		pageRequest := entities.Pageable{}
		// when
		page, err := repo.Search(s.testDb.Ctx, accountId, pageRequest)
		// then
		s.Nil(err)
		s.Len(page.Elements, 0)
	})
}

func (s *RepositoryTestSuite) TestCreateOrder() {
	repo := NewOrderRepository(s.testDb.BunDb)

	s.Run("should create order", func() {
		// given
		order := entities.NewOrderBuilder().
			AccountID(uuid.MustParse("220cea28-b2b0-4051-9eb6-9a99e451af01")).
			Build()

		orderItem1 := entities.NewOrderItemBuilder().
			ItemID(uuid.MustParse("200cea28-b2b0-4051-9eb6-9a99e451af01")).
			OrderID(order.ID).
			Quantity(1).
			Build()

		orderItem2 := entities.NewOrderItemBuilder().
			ItemID(uuid.MustParse("200cea28-b2b0-4051-9eb6-9a99e451af02")).
			OrderID(order.ID).
			Quantity(1).
			Build()

		order.OrderItems = []*entities.OrderItem{orderItem1, orderItem2}

		// when
		err := repo.Create(s.testDb.Ctx, order)
		// then
		s.Nil(err)

		createdOrder, err := repo.GetById(s.testDb.Ctx, order.ID)
		s.Nil(err)
		s.Equal(order.AccountID, createdOrder.AccountID)
		s.NotNil(createdOrder.OrderItems)
		s.Equal(order.OrderItems[0].ItemID, createdOrder.OrderItems[0].ItemID)
	})

	s.Run("should return error if order is nil", func() {
		// given
		var order *entities.Order
		// when
		err := repo.Create(s.testDb.Ctx, order)
		// then
		s.NotNil(err)
	})
}
