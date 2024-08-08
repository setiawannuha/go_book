package routers

import (
	"github.com/setiawannuha/go_book/internal/handlers"
	"github.com/setiawannuha/go_book/internal/middleware"
	"github.com/setiawannuha/go_book/internal/repository"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func bookRouter(g *gin.Engine, d *sqlx.DB) {
	route := g.Group("/book")

	var repo repository.BookRepositoryInterface = repository.NewBookRepository(d)
	handler := handlers.NewBookHandler(repo)

	route.GET("/", middleware.LogMiddleware(), middleware.AuthJwtMiddleware(), handler.GetAll)
	route.GET("/:id", handler.GetDetail)
	route.POST("/", middleware.AuthJwtMiddleware(), handler.Create)
	route.PUT("/:id", middleware.AuthJwtMiddleware(), handler.Update)
	route.DELETE("/:id", middleware.AuthJwtMiddleware(), handler.Delete)
}
