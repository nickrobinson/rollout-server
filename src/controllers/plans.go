package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nickrobinson/rollout-server/src/models"
)

type CreatePlanInput struct {
	Plan models.Plan `json:"plan"`
}

type UpdatePlanInput struct {
	Plan models.Plan `json:"plan"`
}

// GET /plans
// Get all plans
func FindPlans(c *gin.Context) {
	var plans []models.Plan
	models.DB.Find(&plans)

	c.JSON(http.StatusOK, gin.H{"plans": plans})
}

// POST /plans
// Create newplan
func CreatePlan(c *gin.Context) {
	// Validate input
	var input CreatePlanInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// CreatePlan
	plan := models.Plan{Title: input.Plan.Title, Author: input.Plan.Author, Operator: input.Plan.Operator, Overview: input.Plan.Overview, RollbackPlan: input.Plan.RollbackPlan}
	plan.Status = "DRAFT"
	print(plan.RollbackPlan)
	models.DB.Create(&plan)

	c.JSON(http.StatusOK, gin.H{"plans": plan})
}

// GET /plans/:id
// Find a plan
func FindPlan(c *gin.Context) { // Get model if exist
	var plan models.Plan

	if err := models.DB.Where("id = ?", c.Param("id")).First(&plan).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"plans": plan})
}

// PUT /plans/:id
// Update a plan
func UpdatePlan(c *gin.Context) {
	// Get model if exist
	var plan models.Plan
	if err := models.DB.Where("id = ?", c.Param("id")).First(&plan).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input UpdatePlanInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&plan).Updates(input.Plan)

	c.JSON(http.StatusOK, gin.H{"plans": plan})
}

// DELETE /plans/:id
// Delete a plan
func DeletePlan(c *gin.Context) {
	// Get model if exist
	var plan models.Plan
	if err := models.DB.Where("id = ?", c.Param("id")).First(&plan).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&plan)

	c.JSON(http.StatusOK, gin.H{"plans": true})
}
