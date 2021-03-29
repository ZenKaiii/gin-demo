package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK,gin.H{
			"name" : "小哥",
			"desc" : "hhh",
		})
	})
	r.Run(":8889")
}
