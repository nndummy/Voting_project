package middleware

import (
	"log"
	"time"
	"voting_system/app/services"
	"voting_system/app/services/models"

	ginJwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
)

type login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

const JWT_KEY = "secure_as_hell_key"

var identityKey = "id"

func GinJWTMiddlewareHandler() *ginJwt.GinJWTMiddleware {
	authMiddleware, err := ginJwt.New(&ginJwt.GinJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte(JWT_KEY),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: identityKey,
		PayloadFunc: func(data interface{}) ginJwt.MapClaims {
			if v, ok := data.(*models.Voter); ok {
				return ginJwt.MapClaims{
					identityKey: v.StudentId,
				}
			}
			return ginJwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := ginJwt.ExtractClaims(c)
			return &models.Voter{
				StudentId: claims["id"].(string),
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVals login
			if err := c.ShouldBind(&loginVals); err != nil {
				return "", ginJwt.ErrMissingLoginValues
			}
			studentId := loginVals.Username
			password := loginVals.Password

			voter := services.GetVoterInfo(studentId, password)
			if voter != (models.Voter{}) {
				return &models.Voter{
					StudentId: studentId,
					Name:      voter.Name,
				}, nil
			}
			return nil, ginJwt.ErrFailedAuthentication
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {

			if _, ok := data.(*models.Voter); ok {
				return true
			}

			return false

		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		// - "param:<name>"
		TokenLookup: "header: Authorization",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	})
	if err != nil {
		log.Fatal(err)
	}
	return authMiddleware
}
