package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/setiawannuha/go_book/internal/repository"
	"github.com/setiawannuha/go_book/pkg"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRegister(t *testing.T) {
	router := gin.Default()
	userRepositoryMock := new(repository.UserRepositoryMock)
	authRepositoryMock := new(repository.AuthRepositoryMock)

	handler := NewAuthHandler(userRepositoryMock, authRepositoryMock)
	userRepositoryMock.On("CreateUser", mock.Anything).Return("data created success", nil)
	router.POST("/auth/register", handler.Register)

	requestBody, _ := json.Marshal(map[string]string{
		"email":    "testing@gmail.com",
		"password": "123456",
	})
	// membuat request
	req, err := http.NewRequest("POST", "/auth/register", bytes.NewBuffer(requestBody))
	assert.NoError(t, err, "Terjadi error saat membuat request")
	req.Header.Set("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusCreated, recorder.Code, "Status code tidak sesuai")
	var actualResponse pkg.Response
	err = json.Unmarshal(recorder.Body.Bytes(), &actualResponse)
	assert.NoError(t, err, "Terjadi error saat mendapatkan response")

	assert.Equal(t, 201, actualResponse.Status, "Status tidak sesuai")
	assert.Equal(t, "create data success", actualResponse.Message, "Status tidak sesuai")
	assert.Equal(t, "data created success", actualResponse.Data, "Status tidak sesuai")
}
