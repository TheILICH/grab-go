package middleware

import (
	// "go_final/handlers"
	"net/http"

	"github.com/gin-gonic/gin"
	// "github.com/golang-jwt/jwt/v5"
)

func Already() gin.HandlerFunc {
    return func(ctx *gin.Context) {
        // Retrieve the JWT token from the cookie
        token, _ := ctx.Cookie("jwt")
        // if err != nil {
            // ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
            //     "error": "No JWT token found in cookies",
            // })
            // return
        // }

		if token != "" {
			ctx.Redirect(http.StatusFound, "/")
		} else {
			ctx.Next()
		}

    }
}
