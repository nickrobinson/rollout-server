package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nickrobinson/rollout-server/models"
)

type CreatePlanCommentInput struct {
	PlanComment models.PlanComment `json:"planComment"`
}

// GET /planComments
// Get all plans
func FindPlanComments(c *gin.Context) {
	var comments []models.PlanComment
	models.DB.Find(&comments)

	c.JSON(http.StatusOK, gin.H{"planComments": comments})
}

// POST /planComments
// Create new comment
func CreatePlanComment(c *gin.Context) {
	// Validate input
	var input CreatePlanCommentInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// CreatePlanComment
	comment := models.PlanComment{Author: input.PlanComment.Author, PlanID: input.PlanComment.PlanID, Body: input.PlanComment.Body}
	models.DB.Create(&comment)

	c.JSON(http.StatusOK, gin.H{"planComment": comment})
}
