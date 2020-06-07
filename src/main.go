package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nickrobinson/rollout-server/src/controllers"
	"github.com/nickrobinson/rollout-server/src/models"
)

func main() {
	r := gin.Default()

	// Connect to database
	models.ConnectDatabase()

	// Routes
	r.GET("/plans", controllers.FindPlans)
	r.POST("/plans", controllers.CreatePlan)
	r.GET("/plans/:id", controllers.FindPlan)
	r.PATCH("/plans/:id", controllers.UpdatePlan)
	r.DELETE("/plans/:id", controllers.DeletePlan)

	// Run the server
	r.Run()
}
