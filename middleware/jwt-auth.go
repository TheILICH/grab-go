package middleware

import (
	"go_final/handlers"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthorizeJWT() gin.HandlerFunc {
    return func(ctx *gin.Context) {
        // Retrieve the JWT token from the cookie
        tokenString, err := ctx.Cookie("jwt")
        if err != nil {
            // ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
            //     "error": "No JWT token found in cookies",
            // })
            return
        }

        if token, err := handlers.ValidateToken(tokenString); err != nil {
            ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
                "error": "Invalid token",
            })
            return
        } else {
            if claims, ok := token.Claims.(jwt.MapClaims); !ok || !token.Valid {
                ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
                    "error": "Invalid token claims",
                })
                return
            } else {
                ctx.Set("userID", claims["userID"])
                ctx.Set("role", claims["role"])
                ctx.Set("username", claims["username"])
                ctx.Next()
            }
        }
    }
}
