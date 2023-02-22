package main

import (
	"github.com/gin-gonic/gin"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Pong",
		})

	})

	r.POST("/create", func(ctx *gin.Context) {
		var person *Person
		err := ctx.BindJSON(&person)
		if err == nil {
			ctx.AbortWithStatus(500)
		}

		ctx.JSON(200, gin.H{
			"model": person,
		})
	})

	// Start the Server
	r.Run()
}
