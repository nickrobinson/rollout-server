package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nickrobinson/rollout-server/src/models"
)

type CreatePlanInput struct {
	Title     string     `json:"title" binding:"required"`
	Author    string     `json:"author" binding:"required"`
	Operator  string     `json:"operator" binding:"required"`
	StartTime *time.Time `json:start_dt`
	EndTime   *time.Time `json:end_dt binding: "gtfield=StartTime"`
	Overview  string     `json:"overview" binding: "max=1024"`
	Status    string     `json:status`
}

type UpdatePlanInput struct {
	Title     string     `json:"title"`
	Author    string     `json:"author"`
	Operator  string     `json:"operator"`
	StartTime *time.Time `json:start_dt`
	EndTime   *time.Time `json:end_dt binding: "gtfield=StartTime"`
	Overview  string     `json:"overview" binding: "max=1024"`
	Status    string     `json:status`
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
	plan := models.Plan{Title: input.Title, Author: input.Author, Operator: input.Operator, Overview: input.Overview}
	plan.Status = "DRAFT"
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

	models.DB.Model(&plan).Updates(input)

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
