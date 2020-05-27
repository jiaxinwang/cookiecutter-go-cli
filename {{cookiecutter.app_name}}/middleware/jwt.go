package middleware

import (
	"net/http"
	"time"

	"{% if cookiecutter.use_github == "y" -%}github.com/{{cookiecutter.github_username}}/{%- endif %}{{cookiecutter.app_name}}/auth"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

var identityKey = "id"

func helloHandler(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	user, _ := c.Get(identityKey)
	c.JSON(200, gin.H{
		"userID":   claims[identityKey],
		"userName": user.(*User).UserName,
		"text":     "Hello World.",
	})
}

// User demo
type User struct {
	UserName string
}

// AuthMiddleware ...
var AuthMiddleware *jwt.GinJWTMiddleware

// InitJWT ...
func InitJWT() {
	var err error
	// the jwt middleware
	AuthMiddleware, err = jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "major_tom",
		Key:         []byte("_$N(raV@z4QN2IC3Z3U_bms0Zfn4~*9td@6eY7Bn^%OZhAL+J#@3XDxiI~&mxXeYnv5jkEBq~<fxjElmHqnCP>M>JK*cKYUafK6)QqdBBPp>bW8ryCOSCtyZbC|+!:6T"),
		Timeout:     time.Hour * 365 * 24,
		MaxRefresh:  time.Hour,
		IdentityKey: identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*User); ok {
				return jwt.MapClaims{
					identityKey: v.UserName,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &User{
				UserName: claims[identityKey].(string),
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVals login
			if err := c.ShouldBind(&loginVals); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			userID := loginVals.Username
			password := loginVals.Password

			userState := auth.Perm.UserState()
			if !userState.HasUser(userID) {
				return nil, jwt.ErrFailedAuthentication
			}
			if !userState.CorrectPassword(userID, password) {
				return nil, jwt.ErrFailedAuthentication
			} else {
				return &User{UserName: userID}, nil
			}
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			if v, ok := data.(*User); ok {
				userState := auth.Perm.UserState()
				return userState.IsAdmin(v.UserName)
			}
			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			requestID := c.MustGet("requestID")
			c.JSON(200, gin.H{
				"data":       "",
				"error_no":   code,
				"error_msg":  message,
				"request_id": requestID,
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
		TokenLookup: "header: Authorization, query: token, cookie: jwt",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
		LoginResponse: func(c *gin.Context, code int, token string, expire time.Time) {
			requestID := c.MustGet("requestID")
			c.JSON(http.StatusOK, gin.H{
				"data":       gin.H{"token": token, "expire": expire},
				"error_no":   code,
				"error_msg":  "",
				"request_id": requestID,
			})
		},
	})

	if err != nil {
		logrus.WithError(err).Error()
	}
}
