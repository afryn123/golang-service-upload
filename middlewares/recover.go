package middlewares

import (
	"log"
	"net/http"
	"runtime/debug"

	"github.com/gin-gonic/gin"
)

func CustomRecoverPanic() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				stackTrace := debug.Stack()
				log.Printf("[PANIC] %v\n%s\n", r, stackTrace)

				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"status":  false,
					"message": "Internal Server Error",
					"error":   r,
				})
			}
		}()
		c.Next()
	}
}
