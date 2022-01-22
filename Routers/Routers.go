package Routers

import (
	"fmt"

	"github.com/DeniesKresna/beinventaris/Controllers"
	"github.com/DeniesKresna/beinventaris/Middlewares"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.New()
	r.Use(cors.New(cors.Config{
		//AllowOrigins:     []string{"https://foo.com"},
		AllowAllOrigins: true,
		AllowMethods:    []string{"PUT", "PATCH", "GET", "POST", "DELETE", "OPTIONS"},
		AllowHeaders:    []string{"Origin", "X-Requested-With", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:   []string{"Content-Disposition"}, /*
			AllowCredentials: true,
			AllowOriginFunc: func(origin string) bool {
				return origin == "https://github.com"
			},
			MaxAge: 12 * time.Hour,*/
	}))
	// Serve frontend static files
	r.Use(static.Serve("/", static.LocalFile("./Client/Build", true)))
	v1 := r.Group("/api/v1")
	{
		auth := v1.Group("/", Middlewares.Auth("administrator"))

		auth.GET("/users", Controllers.UserIndex)
		auth.GET("/users/me", Controllers.UserMe)
		auth.GET("/users/reset/:id", Controllers.UserReset)
		auth.POST("/users", Controllers.UserStore)
		auth.POST("/users/change-password", Controllers.UserChangePassword)
		auth.PATCH("/users/:id", Controllers.UserUpdate)

		auth.GET("/roles", Controllers.RoleIndex)
		auth.GET("/roles/list", Controllers.RoleList)
		auth.POST("/roles", Controllers.RoleStore)
		auth.PUT("/roles/:id", Controllers.RoleUpdate)

		v1.POST("users/login", Controllers.AuthLogin)

		v1.GET("/medias", func(c *gin.Context) {
			fmt.Print("./Assets/logo.jpeg")
			mediaFile := c.Query("path")
			c.File(mediaFile)
		})

		//v1.GET("users", Controllers.UserIndex)
		//v1.GET("users/:id", Controllers.ShowUser)
		//v1.PUT("users/:id", Controllers.UserUpdate)
		//v1.DELETE("users/:id", Controllers.DestroyUser)
	}
	return r
}
