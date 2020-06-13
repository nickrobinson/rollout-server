package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/nickrobinson/rollout-server/controllers"
	"github.com/nickrobinson/rollout-server/models"
	viper "github.com/spf13/viper"
)

var identityKey = "id"

// User demo
type User struct {
	UserName string
}

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

	// the jwt middleware
	authMiddleware, _ := jwt.New(&jwt.GinJWTMiddleware{
		Realm:            "rollout",
		SigningAlgorithm: "HS256",
		Key:              []byte(viper.GetString("app.server.jwtSecretKey")),
		Timeout:          time.Hour,
		MaxRefresh:       time.Hour,
		IdentityKey:      identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			return &User{}
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

	if viper.GetBool("app.server.disableAuth") == false {
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

	// Run the server
	r.Run()
}
