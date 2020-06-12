package main

import (
	"os"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/nickrobinson/rollout-server/controllers"
	"github.com/nickrobinson/rollout-server/models"
)

var identityKey = "id"

type login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

// User demo
type User struct {
	UserName  string
	FirstName string
	LastName  string
}

func main() {
	r := gin.Default()

	// the jwt middleware
	authMiddleware, _ := jwt.New(&jwt.GinJWTMiddleware{
		Realm:            "test zone",
		SigningAlgorithm: "HS256",
		Key:              []byte(os.Getenv("AUTH_SECRET")),
		Timeout:          time.Hour,
		MaxRefresh:       time.Hour,
		IdentityKey:      identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			return &User{
				UserName: "nick",
			}
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			return true
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},

		TokenLookup: "header: Authorization, cookie: jwt",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	})

	r.Use(authMiddleware.MiddlewareFunc())

	// Connect to database
	models.ConnectDatabase()

	// Routes
	r.GET("/plans", controllers.FindPlans)
	r.POST("/plans", controllers.CreatePlan)
	r.GET("/plans/:id", controllers.FindPlan)
	r.PUT("/plans/:id", controllers.UpdatePlan)
	r.DELETE("/plans/:id", controllers.DeletePlan)

	// Run the server
	r.Run()
}
