package handlers

import (
	"github.com/setiawannuha/go_book/internal/models"
	"github.com/setiawannuha/go_book/internal/repository"
	"github.com/setiawannuha/go_book/pkg"

	"github.com/gin-gonic/gin"
)

type BookHandler struct {
	repository.BookRepositoryInterface
}

func NewBookHandler(r repository.BookRepositoryInterface) *BookHandler {
	return &BookHandler{r}
}

func (h *BookHandler) Create(ctx *gin.Context) {
	response := pkg.NewResponse(ctx)
	body := models.Book{}
	if err := ctx.ShouldBind(&body); err != nil {
		response.BadRequest("create data failed", err.Error())
		return
	}
	result, err := h.CreateData(&body)
	if err != nil {
		response.BadRequest("create data failed", err.Error())
		return
	}
	response.Created("create data success", result)
}

func (h *BookHandler) GetDetail(ctx *gin.Context) {
	response := pkg.NewResponse(ctx)
	id := ctx.Param("id")
	result, err := h.GetDetailData(id)
	if err != nil {
		response.InternalServerError("get data failed", err.Error())
		return
	}
	response.Success("get data success", result)
}

func (h *BookHandler) GetAll(ctx *gin.Context) {
	response := pkg.NewResponse(ctx)
	result, err := h.GetAllData()
	if err != nil {
		response.InternalServerError("get data failed", err.Error())
		return
	}
	response.Success("get data success", result)
}

func (h *BookHandler) Update(ctx *gin.Context) {
	response := pkg.NewResponse(ctx)
	id := ctx.Param("id")
	body := models.Book{}
	if err := ctx.ShouldBind(&body); err != nil {
		response.BadRequest("update data failed", err.Error())
		return
	}
	result, err := h.UpdateData(id, &body)
	if err != nil {
		response.BadRequest("update data failed", err.Error())
		return
	}
	response.Success("update data success", result)
}

func (h *BookHandler) Delete(ctx *gin.Context) {
	response := pkg.NewResponse(ctx)
	id := ctx.Param("id")
	result, err := h.DeleteData(id)
	if err != nil {
		response.BadRequest("delete data failed", err.Error())
		return
	}
	response.Success("delete data success", result)
}
