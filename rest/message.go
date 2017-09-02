package rest

import (
	"BooBooMessenger-server/model"
	"BooBooMessenger-server/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

// MessageController .
type MessageController struct {
	r *repository.MessageRepository
}

// NewMessageController .
func NewMessageController(repository *repository.MessageRepository) *MessageController {
	return &MessageController{r: repository}
}

// GetAllMessage .
func (controller *MessageController) GetAllMessage(c *gin.Context) {
	messages, err := controller.r.GetAllMessages()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		panic(err)
	}

	c.JSON(200, messages)
}

// MarkMessageAsRead .
func (controller *MessageController) MarkMessageAsRead(c *gin.Context) {
	id := c.Param("id")
	err := controller.r.MarkMessageAsRead(id)

	if err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		panic(err)
	}

}

// SaveNewMessage .
func (controller *MessageController) SaveNewMessage(c *gin.Context) {
	var message model.Message

	err := c.BindJSON(&message)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		panic(err)
	}

	savedMessage, err := controller.r.SaveMessage(&message)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		panic(err)
	}

	c.JSON(http.StatusCreated, savedMessage)
}