package rest

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vincentaudoire/BooBooMessenger-server/repository"
)

type PrinterController struct {
	r *repository.PrinterRepository
}

func NewPrinterController(r *repository.PrinterRepository) *PrinterController {
	return &PrinterController{r}
}

func (controller *PrinterController) RegisterNewPrinter(c *gin.Context) {

	var printer struct {
		UUID string `json:"uuid"`
		Name string `json:"name"`
	}

	err := c.BindJSON(&printer)

	if err != nil {
		log.Panic(err)
		return
	}

	log.Println("Registering new printer: " + printer.Name + " " + printer.UUID)

	err = controller.r.RegisterNewPrinter(printer.UUID, printer.Name)

	if err != nil {
		log.Panic(err)
		return
	}

	c.JSON(http.StatusOK, nil)
}

func (controller *PrinterController) GetAllMessages(c *gin.Context) {
	printerID := "1" //TODO: Real printer ID
	messages, err := controller.r.GetAllPendingPrintingMessages(printerID)
	if err != nil {
		log.Println(err.Error())
	}

	c.JSON(http.StatusOK, messages)
}
