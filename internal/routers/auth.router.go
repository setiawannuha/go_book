package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/setiawannuha/go_book/internal/handlers"
	"github.com/setiawannuha/go_book/internal/repository"
)

func authRouter(g *gin.Engine, d *sqlx.DB) {
	route := g.Group("/auth")

	var userRepo repository.UserRepositoryInterface = repository.NewUserRepository(d)
	var authRepo repository.AuthRepositoryInterface = repository.NewAuthRepository(d)
	handler := handlers.NewAuthHandler(userRepo, authRepo)

	route.POST("/register", handler.Register)
	route.POST("/login", handler.Login)
}
