package api

import (
	"github.com/arman-ala/Car_shop/api/routers"
	"github.com/gin-gonic/gin"
)

func InitServer() {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	// api version
	v1 := r.Group("/api/v1/")
	{
		healthGroup := v1.Group("/health")

		routers.HealthRouter(healthGroup)
	}

	r.Run(":9000")
}
