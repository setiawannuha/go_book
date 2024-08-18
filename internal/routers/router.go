package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func NewRouter(db *sqlx.DB) *gin.Engine {
	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		html := `<html><body><h1>It Works!</h1></body></html>`
		ctx.Data(200, "text/html; charset=utf-8", []byte(html))
	})
	bookRouter(router, db)
	authRouter(router, db)

	return router
}
