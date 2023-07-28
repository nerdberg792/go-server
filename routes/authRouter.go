package routes

import (
	controller "github.com/nerdber792/go-server/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("users/signup", controller.signUp())
	incomingRoutes.POST("users/login", controller.Login)

}
