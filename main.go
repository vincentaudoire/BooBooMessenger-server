package main

import (
	"BooBooMessenger-server/model"
	"fmt"

	"github.com/gin-gonic/gin"
)

var messages []model.Message

func main() {
	router := gin.Default()

	router.GET("/messages", getMessages)
	router.POST("/messages", postMessage)

	router.Run() // listen and serve on 0.0.0.0:8080
}

func getMessages(c *gin.Context) {
	c.JSON(200, messages)
}

func postMessage(c *gin.Context) {
	var message model.Message

	err := c.BindJSON(&message)

	if err != nil {
		fmt.Printf(err.Error())
	}

	fmt.Printf("message received: %s\n", message.Message)
}
