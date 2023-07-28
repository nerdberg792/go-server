package routes

import (
	"github.com/nerdber792/go-server/middleware"

	controller "github.com/nerdber792/go-server/controllers"

	"github.com/gin-gonic/gin"
)

func userRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users", controller.GetUSers())
	incomingRoutes.GET("/users/:user_id", controller.GetUsers())
	incomingRoutes.PUT("/api/user/:user_id", controller.UpdateUser())
	incomingRoutes.DELETE("/api/user/:user_id", controller.DeleteUser())

}
