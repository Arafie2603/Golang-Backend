// router.go

package router

import (
	"finpro-golang2/controllers"
	"finpro-golang2/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Rute untuk pengguna
	userGroup := r.Group("/users")
	{
		userGroup.POST("/register", controllers.Register)
		userGroup.POST("/login", controllers.Login)

		// Menggunakan otentikasi JWT untuk melindungi rute-rute di bawah ini
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

		// Menggunakan otentikasi JWT untuk melindungi rute-rute di bawah ini
		photoGroup.Use(middlewares.AuthMiddleware())
		{
			photoGroup.PUT("/:photoId", controllers.UpdatePhoto)
			photoGroup.DELETE("/:photoId", controllers.DeletePhoto)
		}
	}

	return r
}