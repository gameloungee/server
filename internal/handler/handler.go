package handler

import "github.com/gin-gonic/gin"

func InitHanlder() *gin.Engine {
	r := gin.New()

	v1 := r.Group("/api/v1/")
	{
		acc := v1.Group("account/")
		{
			acc.POST("create", CreateAccountHandler)
		}
	}

	return r
}
