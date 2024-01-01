// router.go

package router

import (
	"finpro-golang2/controllers"
	middlewares "finpro-golang2/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	userGroup := r.Group("/users")
	{
		userGroup.POST("/register", controllers.Register)
		userGroup.POST("/login", controllers.Login)

		userGroup.Use(middlewares.AuthMiddleware())
		{
			userGroup.PUT("/:userId", controllers.UpdateUser)
			userGroup.DELETE("/:userId", controllers.DeleteUser)
		}
	}

	// Rute untuk foto
	photoGroup := r.Group("/photos")
	{
		photoGroup.POST("/createPhoto", controllers.CreatePhoto)
		photoGroup.GET("", controllers.GetPhotos)

		photoGroup.Use(middlewares.AuthMiddleware())
		{
			photoGroup.PUT("/:photoId", controllers.UpdatePhoto)
			photoGroup.DELETE("/:photoId", controllers.DeletePhoto)
		}
	}

	return r
}
