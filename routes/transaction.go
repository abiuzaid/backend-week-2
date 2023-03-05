package routes

import (
	"backEnd/handlers"
	"backEnd/pkg/middleware"
	"backEnd/pkg/mysql"
	"backEnd/repositories"

	"github.com/labstack/echo/v4"
)

func TransactionRoutes(e *echo.Group) {
	transactionRepository := repositories.RepositoryTransaction(mysql.DB)
	h := handlers.HandlerTransaction(transactionRepository)

	e.GET("/transactions", middleware.Auth(h.FindTransactions))
	e.GET("/transaction/:id", middleware.Auth(h.GetTransaction))
	e.POST("/transaction/:product_id", middleware.Auth(h.CreateTransaction))
	e.PATCH("/transaction/:id", middleware.Auth(h.UpdateTransaction))
	e.DELETE("/transaction/:id", middleware.Auth(h.DeleteTransaction))
}
