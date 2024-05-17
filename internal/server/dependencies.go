package server

import (
	"github.com/fmiskovic/new-amz/internal/core/dtos"
	"github.com/fmiskovic/new-amz/internal/core/entities"
	"github.com/fmiskovic/new-amz/internal/core/services"
	"github.com/fmiskovic/new-amz/internal/db"
	"github.com/fmiskovic/new-amz/internal/handlers"
	"github.com/fmiskovic/new-amz/internal/handlers/mappers"
	"github.com/fmiskovic/new-amz/internal/repositories"
	"github.com/google/uuid"
)

// This struct is used to wire dependencies with server.
// It is created by the bootstrap function.
// Add any new dependencies here.
type dependencies struct {
	// handlers
	createAccountHandler       handlers.Handler[dtos.CreateAccountCommand, dtos.CreateAccountAnswer]
	getAccountByIdHandler      handlers.Handler[uuid.UUID, dtos.AccountDto]
	getItemByIdHandler         handlers.Handler[uuid.UUID, dtos.ItemDto]
	getItemsPageHandler        handlers.Handler[entities.Pageable, entities.Page[dtos.ItemDto]]
	createOrderHandler         handlers.Handler[dtos.CreateOrderCommand, dtos.CreateOrderAnswer]
	getOrderByIdHandler        handlers.Handler[uuid.UUID, dtos.OrderDto]
	searchAccountOrdersHandler handlers.Handler[dtos.OrderFilter, entities.Page[dtos.OrderDto]]
}

// bootstrap creates and wires up all dependencies.
func bootstrap() dependencies {
	dbSvc := db.NewService()
	sqlDb, err := dbSvc.Connect()
	if err != nil {
		panic(err)
	}
	bunDb := dbSvc.WrapWithBun(sqlDb)

	// Account
	accountRepository := repositories.NewAccountRepository(bunDb)
	accountService := services.NewAccountService(accountRepository)
	createAccountHandler := handlers.New(
		mappers.NewCreateAccountRequestMapper(),
		mappers.NewCreateAccountResponseMapper(),
		accountService.Create,
	)
	getAccountByIdHandler := handlers.New(
		mappers.NewGetAccountByIdRequestMapper(),
		mappers.NewGetAccountByIdResponseMapper(),
		accountService.GetById,
	)

	// Item
	itemRepository := repositories.NewItemRepository(bunDb)
	itemService := services.NewItemService(itemRepository)
	getItemByIdHandler := handlers.New(
		mappers.NewItemGetByIdRequestMapper(),
		mappers.NewItemGetByIdResponseMapper(),
		itemService.GetById,
	)
	getItemsPageHandler := handlers.New(
		mappers.NewItemGetPageRequestMapper(),
		mappers.NewItemGetPageResponseMapper(),
		itemService.GetPage,
	)

	// Order
	orderRepository := repositories.NewOrderRepository(bunDb)
	orderService := services.NewOrderService(orderRepository)
	createOrderHandler := handlers.New(
		mappers.NewOrderCreateRequestMapper(),
		mappers.NewOrderCreateResponseMapper(),
		orderService.Create,
	)
	getOrderByIdHandler := handlers.New(
		mappers.NewOrderGetByIdRequestMapper(),
		mappers.NewOrderGetByIdResponseMapper(),
		orderService.GetById,
	)
	searchAccountOrdersHandler := handlers.New(
		mappers.NewOrderSearchRequestMapper(),
		mappers.NewOrderSearchResponseMapper(),
		orderService.Search,
	)

	return dependencies{
		createAccountHandler:       createAccountHandler,
		getAccountByIdHandler:      getAccountByIdHandler,
		getItemByIdHandler:         getItemByIdHandler,
		getItemsPageHandler:        getItemsPageHandler,
		createOrderHandler:         createOrderHandler,
		getOrderByIdHandler:        getOrderByIdHandler,
		searchAccountOrdersHandler: searchAccountOrdersHandler,
	}
}
