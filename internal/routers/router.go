package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func NewRouter(db *sqlx.DB) *gin.Engine {
	router := gin.Default()

	bookRouter(router, db)
	authRouter(router, db)

	return router
}
