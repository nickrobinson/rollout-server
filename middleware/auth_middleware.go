package middleware

import (
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var identityKey = "id"

// User demo
type User struct {
	UserName string
}

func GetAuthMiddleware() (*jwt.GinJWTMiddleware, error) {
	return jwt.New(&jwt.GinJWTMiddleware{
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
}
