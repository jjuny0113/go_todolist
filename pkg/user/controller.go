package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	service Service
}

func NewController(service Service) *Controller {
	return &Controller{service: service}
}

func (c *Controller) Get(ctx *gin.Context) {
	fmt.Println("get!!")

}
