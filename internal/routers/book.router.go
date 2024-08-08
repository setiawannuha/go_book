package routers

import (
	"github.com/setiawannuha/go_book/internal/handlers"
	"github.com/setiawannuha/go_book/internal/repository"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func bookRouter(g *gin.Engine, d *sqlx.DB) {
	route := g.Group("/book")

	var repo repository.BookRepositoryInterface = repository.NewBookRepository(d)
	handler := handlers.NewBookHandler(repo)

	route.GET("/", handler.GetAll)
	route.GET("/:id", handler.GetDetail)
	route.POST("/", handler.Create)
	route.PUT("/:id", handler.Update)
	route.DELETE("/:id", handler.Delete)
}
