package main

import (
	"errors"

	"github.com/gin-gonic/gin"
)

func foo(bar string) error {
	if len(bar) < 3 {
		return errors.New("ddddd")
	}
	return nil
}

func main() {
	var bar = "test"
	if err := foo(bar); err != nil {
		panic(err)
	}

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() //listen and serve on 0.0.0.0:8080
}
