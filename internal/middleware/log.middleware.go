package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func LogMiddleware() gin.HandlerFunc {
	return func(cxt *gin.Context) {

		fmt.Println("melewati middleware")
		cxt.Next()
	}
}
