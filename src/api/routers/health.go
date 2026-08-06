package routers

import (
	"github.com/arman-ala/Car_shop/api/handlers"
	"github.com/gin-gonic/gin"
)

func HealthRouter(r *gin.RouterGroup) {
	handler := handlers.NewHealthHandler()
	r.GET("/", handler.Health)
}
