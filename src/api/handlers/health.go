package handlers

import (
	"github.com/arman-ala/Car_shop/api/helper"
	"github.com/gin-gonic/gin"
)

type HealthHandler struct {
}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) Health(c *gin.Context) {
	c.JSON(200, helper.GenerateBaseResponse(gin.H{
		"message": "/health works fine.",
	}, true,
		0,
	))
}
