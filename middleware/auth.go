package middleware

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		sessionID := session.Get("session_token")
		if sessionID == nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": "unauthorized",
				"error":  "session habis, silahkan login kembali",
			})
			c.Abort()
		}
	}
}
