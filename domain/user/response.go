package user

import "github.com/gin-gonic/gin"

type responseError struct {
	Success   bool   `json:"success"`
	ErrorCode int    `json:"error_code"`
	Message   string `json:"message"`
}

type responseSuccess struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
}

func writeSuccess(ctx *gin.Context, data interface{}, statusCode int) {
	resp := responseSuccess{
		Success: true,
		Data:    data,
	}

	ctx.JSON(statusCode, resp)
}

func writeError(ctx *gin.Context, message string, errorCode int, statusCode int) {
	resp := responseError{
		Success:   false,
		ErrorCode: errorCode,
		Message:   message,
	}

	ctx.JSON(statusCode, resp)
}
