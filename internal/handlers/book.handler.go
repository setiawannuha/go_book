package handlers

import (
	"fmt"
	"math/rand"

	"github.com/asaskevich/govalidator"
	"github.com/setiawannuha/go_book/internal/models"
	"github.com/setiawannuha/go_book/internal/repository"
	"github.com/setiawannuha/go_book/pkg"

	"github.com/gin-gonic/gin"
)

type BookHandler struct {
	repository.BookRepositoryInterface
	pkg.Cloudinary
}

func NewBookHandler(r repository.BookRepositoryInterface, cld pkg.Cloudinary) *BookHandler {
	return &BookHandler{r, cld}
}

func (h *BookHandler) Create(ctx *gin.Context) {
	response := pkg.NewResponse(ctx)
	body := models.Book{}
	if err := ctx.ShouldBind(&body); err != nil {
		response.BadRequest("create data failed", err.Error())
		return
	}
	_, err := govalidator.ValidateStruct(&body)
	if err != nil {
		response.BadRequest("create data failed", err.Error())
		return
	}

	// get file from request body
	file, header, err := ctx.Request.FormFile("image")
	if err != nil {
		response.BadRequest("create data failed, upload file failed", err.Error())
		return
	}
	// validate file ext
	fmt.Println(header.Size)
	mimeType := header.Header.Get("Content-Type")
	if mimeType != "image/jpg" && mimeType != "image/png" {
		response.BadRequest("create data failed, upload file failed, wrong file type", err.Error())
		return
	}
	// upload file
	randomNumber := rand.Int()
	fileName := fmt.Sprintf("go-book-%d", randomNumber)
	uploadResult, err := h.UploadFile(ctx, file, fileName)
	if err != nil {
		response.BadRequest("create data failed, upload file failed", err.Error())
		return
	}
	body.Image = uploadResult.SecureURL

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
	fmt.Println(ctx.Get("userId"))
	fmt.Println(ctx.Get("email"))
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
