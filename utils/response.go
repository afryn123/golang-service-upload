package utils

import (
	"github.com/gin-gonic/gin"
)

type APIResponse struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Error   interface{} `json:"error"`
}

func JSONResponse(c *gin.Context, status int, message string, data interface{}) {
	c.JSON(status, APIResponse{
		Status:  true,
		Message: message,
		Data:    data,
	})
}

func ErrorResponse(c *gin.Context, status int, message string, error interface{}) {
	c.JSON(status, APIResponse{
		Status:  false,
		Message: message,
		Error:   error,
	})
}
