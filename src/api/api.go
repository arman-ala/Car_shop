package api

import (
	"fmt"
	"log"

	"github.com/arman-ala/Car_shop/api/middlewares"
	"github.com/arman-ala/Car_shop/api/routers"
	"github.com/arman-ala/Car_shop/api/validations"
	"github.com/arman-ala/Car_shop/config"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func InitServer() {
	cfg := config.SetConfig()
	r := gin.New()

	// register custom validators
	validator, ok := binding.Validator.Engine().(*validator.Validate)
	if !ok {
		log.Printf("error happened while initializing validator: validator not found")
		return
	}
	validator.RegisterValidation("IR_phone_number", validations.IranianPhoneNumberValidator, true)
	// TODO: I should use config to validate password
	validator.RegisterValidation("password", validations.PasswordValidator, true)

	r.Use(gin.Logger(), gin.Recovery(), middlewares.LimitByRequestMiddleware())
	// api group
	api := r.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			healthGroup := v1.Group("/health")

			routers.HealthRouter(healthGroup)
		}
	}

	r.Run(fmt.Sprintf(":%s", cfg.Server.Port))
}
