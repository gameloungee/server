package response

import "github.com/gin-gonic/gin"

const DETAILS_HIDDEN = "Details not available due to potential danger to system security"

type Response struct {
	Code    int         `json:"code"`
	Message interface{} `json:"message"`
	Details interface{} `json:"details"`
}

func Send(statusCode int, msg, details interface{}, c *gin.Context) {
	c.JSON(statusCode, Response{
		Code:    statusCode,
		Message: msg,
		Details: details,
	})
}

func AbortWith(statusCode int, msg, details interface{}, c *gin.Context) {
	c.AbortWithStatusJSON(statusCode, Response{
		Code:    statusCode,
		Message: msg,
		Details: details,
	})
}
