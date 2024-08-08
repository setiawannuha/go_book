package handlers

import (
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/setiawannuha/go_book/internal/models"
	"github.com/setiawannuha/go_book/internal/repository"
	"github.com/setiawannuha/go_book/pkg"
)

type AuthHandler struct {
	repository.UserRepositoryInterface
	repository.AuthRepositoryInterface
}

func NewAuthHandler(userRepo repository.UserRepositoryInterface, authRepo repository.AuthRepositoryInterface) *AuthHandler {
	return &AuthHandler{userRepo, authRepo}
}

func (h *AuthHandler) Register(ctx *gin.Context) {
	response := pkg.NewResponse(ctx)
	body := models.User{}

	if err := ctx.ShouldBind(&body); err != nil {
		response.BadRequest("create data failed", err.Error())
		return
	}
	_, err := govalidator.ValidateStruct(&body)
	if err != nil {
		response.BadRequest("create data failed", err.Error())
		return
	}

	body.Password, err = pkg.HashPassword(body.Password)
	if err != nil {
		response.BadRequest("create data failed", err.Error())
		return
	}

	result, err := h.CreateUser(&body)
	if err != nil {
		response.BadRequest("create data failed", err.Error())
		return
	}

	response.Created("create data success", result)
}

func (h *AuthHandler) Login(ctx *gin.Context) {
	response := pkg.NewResponse(ctx)
	body := models.User{}

	if err := ctx.ShouldBind(&body); err != nil {
		response.BadRequest("login failed", err.Error())
		return
	}
	_, err := govalidator.ValidateStruct(&body)
	if err != nil {
		response.BadRequest("login failed", err.Error())
		return
	}

	result, err := h.GetByEmail(body.Email)
	if err != nil {
		response.BadRequest("login failed", err.Error())
		return
	}

	err = pkg.VerifyPassword(result.Password, body.Password)
	if err != nil {
		response.Unauthorized("wrong password", err.Error())
		return
	}

	jwt := pkg.NewJWT(result.Id, result.Email)
	token, err := jwt.GenerateToken()
	if err != nil {
		response.Unauthorized("failed generate token", err.Error())
		return
	}

	response.Created("login success", token)
}
