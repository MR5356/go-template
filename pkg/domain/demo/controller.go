package demo

import (
	"github.com/MR5356/go-template/pkg/response"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	service *Service
}

func NewController(service *Service) *Controller {
	return &Controller{
		service: service,
	}
}

func (c *Controller) handleListDemo(ctx *gin.Context) {
	res, err := c.service.db.List(&Demo{})
	if err != nil {
		response.Error(ctx, response.CodeClient, "Failed to list demo")
	} else {
		response.Success(ctx, res)
	}
}

func (c *Controller) RegisterRoute(group *gin.RouterGroup) {
	api := group.Group("/demo")
	api.GET("/list", c.handleListDemo)
}
