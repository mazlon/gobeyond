package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mazlon/gobeyond/internal/auth"
)

func AuthorizationMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie(`token`)
		if err != nil {
			c.String(http.StatusUnauthorized, "Unauthorized, no token found")
			c.Abort()
			return
		}
		uID, err := auth.TokenVerifier(token)
		if err != nil {
			c.String(http.StatusUnauthorized, "Unauthorized, invalid token")
			c.Abort()
			return
		}
		if uID == "" {
			c.String(http.StatusUnauthorized, "Unauthorized, invalid user id")
			c.Abort()
			return
		}
		c.Set("userID", uID)
		c.Next()
	}
}
