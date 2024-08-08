package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/setiawannuha/go_book/pkg"
)

func AuthJwtMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		response := pkg.NewResponse(ctx)
		var header string

		if header = ctx.GetHeader("Authorization"); header == "" {
			response.Unauthorized("Unauthorized", nil)
			return
		}

		if !strings.Contains(header, "Bearer") {
			response.Unauthorized("Inavlid Bearer Token", nil)
			return
		}

		// Bearer Bearer token
		token := strings.Replace(header, "Bearer ", "", -1)

		check, err := pkg.VerifyToken(token)
		if err != nil {
			response.Unauthorized("Inavlid Bearer Token", nil)
			return
		}

		ctx.Set("userId", check.Id)
		ctx.Set("email", check.Email)
		ctx.Next()
	}
}

// func AadminAuthMiddleware() gin.HandlerFunc {
// 	return func(ctx *gin.Context) {
// 		response := pkg.NewResponse(ctx)
// 		var header string

// 		if header = ctx.GetHeader("Authorization"); header == "" {
// 			response.Unauthorized("Unauthorized", nil)
// 			return
// 		}

// 		if !strings.Contains(header, "Bearer") {
// 			response.Unauthorized("Inavlid Bearer Token", nil)
// 			return
// 		}

// 		// Bearer Bearer token
// 		token := strings.Replace(header, "Bearer ", "", -1)

// 		check, err := pkg.VerifyToken(token)
// 		if err != nil {
// 			response.Unauthorized("Inavlid Bearer Token", nil)
// 			return
// 		}

// 		if check.Role != "Admin" {
// 			response.Unauthorized("Inavlid Bearer Token", nil)
// 			return
// 		}
// 		ctx.Next()
// 	}
// }
