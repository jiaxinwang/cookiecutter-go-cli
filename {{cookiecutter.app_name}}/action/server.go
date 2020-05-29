package action

import (
	"github.com/gin-gonic/gin"
	"github.com/urfave/cli/v2"
	"{% if cookiecutter.use_github == "y" -%}github.com/{{cookiecutter.github_username}}/{%- endif %}{{cookiecutter.app_name}}/middleware"
	"{% if cookiecutter.use_github == "y" -%}github.com/{{cookiecutter.github_username}}/{%- endif %}{{cookiecutter.app_name}}/config"
	v1 "{% if cookiecutter.use_github == "y" -%}github.com/{{cookiecutter.github_username}}/{%- endif %}{{cookiecutter.app_name}}/controller/v1"

)

// Server ...
func Server(c *cli.Context) error {
	config.Load(c.String(`conf`))
	GinEngine().Run(config.Config.Server.Listen)
	return nil
}

// GinEngine ...
func GinEngine() *gin.Engine {
	middleware.InitJWT()
	var r *gin.Engine
	r = gin.Default()
	r.Use(middleware.Access)
	r.Use(middleware.Recovery)
	r.Use(middleware.RequestLogger)
	r.GET("/health")
	V1(r)
	return r
}

// V1 api set
func V1(r *gin.Engine) {
	g := r.Group("/v1")
	{
		g.GET("/echo", nil)
		g.POST("/register", v1.Register)
	}
}

// Debug api set
func Debug(r *gin.Engine) {
	r.Use(middleware.AuthMiddleware.MiddlewareFunc())
	{
		r.GET("/hello", nil)
	}
}
