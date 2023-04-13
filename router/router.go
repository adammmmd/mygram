package router

import (
	"Project/middlewares"
	"Project/controllers"
	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", controllers.UserRegister)
		userRouter.POST("/login", controllers.UserLogin)
	}

	photoRouter := r.Group("/photos")
	{
		photoRouter.Use(middlewares.Authentication())
		photoRouter.GET("/view", controllers.GetAllPhoto)
		photoRouter.GET("/:photoId", controllers.GetOnePhoto)
		photoRouter.POST("/", controllers.CreatePhoto)
		photoRouter.PUT("/:photoId", middlewares.PhotoAuthorization(), controllers.UpdatePhoto)
		photoRouter.DELETE("/:photoId", middlewares.PhotoAuthorization(), controllers.DeletePhoto)
	}

	commentRouter := r.Group("/comments")
	{
		commentRouter.Use(middlewares.Authentication())
		commentRouter.GET("/view", controllers.GetAllComment)
		commentRouter.GET("/:commentId", controllers.GetOneComment)
		commentRouter.POST("/", controllers.CreateComment)
		commentRouter.PUT("/:commentId", middlewares.CommentAuthorization(), controllers.UpdateComment)
		commentRouter.DELETE("/:commentId", middlewares.CommentAuthorization(), controllers.DeleteComment)
	}

	socmedRouter := r.Group("/socmedias")
	{
		socmedRouter.Use(middlewares.Authentication())
		socmedRouter.GET("/view", controllers.GetAllSocmed)
		socmedRouter.GET("/:socmedId", controllers.GetOneSocmed)
		socmedRouter.POST("/", controllers.CreateSocmed)
		socmedRouter.PUT("/:socmedId", middlewares.SocmedAuthorization(), controllers.UpdateSocmed)
		socmedRouter.DELETE("/:socmedId", middlewares.SocmedAuthorization(), controllers.DeleteSocmed)
	}

	return r
}