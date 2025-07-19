package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func MethodNotAllowedMiddleware(routes []gin.RouteInfo) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		path := ctx.Request.URL.Path
		method := ctx.Request.Method
		allowedMethods := []string{}
		for _, route := range routes {
			if route.Path == path {
				if route.Method == method {
					ctx.Next()
					return
				}
				allowedMethods = append(allowedMethods, method)
			}
		}
		if len(allowedMethods) > 0 {
			ctx.Header("Allow", strings.Join(allowedMethods, ", "))
			ctx.AbortWithStatus(http.StatusMethodNotAllowed)
			return
		}
		ctx.Next()
	}
}
