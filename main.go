package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/nickrobinson/rollout-server/controllers"
	"github.com/nickrobinson/rollout-server/middleware"
	"github.com/nickrobinson/rollout-server/models"
	viper "github.com/spf13/viper"
)

func main() {
	r := gin.Default()

	// Setup static config
	viper.SetConfigName("default.yml")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	viper.ReadInConfig()

	env := "dev"
	if envVar := os.Getenv("ENV"); envVar != "" {
		env = strings.ToLower(envVar)
	}

	viper.SetConfigName(fmt.Sprintf("%s.yml", env))
	viper.MergeInConfig()

	if viper.GetBool("app.server.disableAuth") == false {
		// the jwt middleware
		authMiddleware, err := middleware.GetAuthMiddleware()
		if err != nil {
			panic("Unable to initialize auth Middleware")
		}
		r.Use(authMiddleware.MiddlewareFunc())
	}

	// Connect to database
	models.ConnectDatabase()

	// Routes
	r.GET("/plans", controllers.FindPlans)
	r.POST("/plans", controllers.CreatePlan)
	r.GET("/plans/:id", controllers.FindPlan)
	r.PUT("/plans/:id", controllers.UpdatePlan)
	r.DELETE("/plans/:id", controllers.DeletePlan)

	r.GET("planComments", controllers.FindPlanComments)
	r.POST("/planComments", controllers.CreatePlanComment)

	// Run the server
	r.Run()
}
