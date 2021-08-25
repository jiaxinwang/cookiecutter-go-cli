package action

import (
	"github.com/gin-gonic/gin"
	"github.com/urfave/cli/v2"
	"{% if cookiecutter.use_github == "y" -%}github.com/{{cookiecutter.github_username}}/{%- endif %}{{cookiecutter.app_name}}/middleware"
	"{% if cookiecutter.use_github == "y" -%}github.com/{{cookiecutter.github_username}}/{%- endif %}{{cookiecutter.app_name}}/auth"
	"{% if cookiecutter.use_github == "y" -%}github.com/{{cookiecutter.github_username}}/{%- endif %}{{cookiecutter.app_name}}/config"
	"{% if cookiecutter.use_github == "y" -%}github.com/{{cookiecutter.github_username}}/{%- endif %}{{cookiecutter.app_name}}/db"
	"{% if cookiecutter.use_github == "y" -%}github.com/{{cookiecutter.github_username}}/{%- endif %}{{cookiecutter.app_name}}/model"
	v1 "{% if cookiecutter.use_github == "y" -%}github.com/{{cookiecutter.github_username}}/{%- endif %}{{cookiecutter.app_name}}/controller/v1"
	gm "github.com/jiaxinwang/common/gin-middleware"
	"github.com/jiaxinwang/lazy"
	idiocy "github.com/jiaxinwang/go-idiocy/doc"
	
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Server ...
func Server(c *cli.Context) error {
	Prepare(c)
	idiocy.Analyse("__PATH_TO_GO_MOD_FILE__")
	doc.Init()
	GinEngine().Run(config.Config.Server.Listen)
	return nil
}

// GinEngine ...
func GinEngine() *gin.Engine {
	auth.Init(config.Config.Permission.DSN, config.Config.Permission.DB)
	middleware.InitJWT()
	r := gin.Default()
	r.Use(middleware.Access)
	r.Use(middleware.Recovery)
	r.Use(middleware.RequestLogger)
	r.GET("/health")
	V1(r)
	DogAPI(r)
	CatAPI(r)
	return r
}

// V1 api set
func V1(r *gin.Engine) {
	g := r.Group("/v1")
	{
		g.GET("/echo", nil)
		g.POST("/signup", v1.Signup)
	}
}

// Debug api set
func Debug(r *gin.Engine) {
	r.Use(middleware.AuthMiddleware.MiddlewareFunc())
	{
		r.GET("/hello", nil)
	}
}

// DogAPI api set
func DogAPI(r *gin.Engine) {
	g := r.Group("/v1/dog").Use(gm.Trace, lazy.MiddlewareParams, lazy.MiddlewareResponse, lazy.MiddlewareDefaultResult, lazy.MiddlewareExec)
	{
		g.GET("", func(c *gin.Context) {
			c.Set(lazy.KeyConfig, &lazy.Configuration{
				DB: db.DB, Model: &model.Dog{}, Action: []lazy.Action{{Action: lazy.DefaultGetAction}}
			})
			return
		})
		g.POST("", func(c *gin.Context) {
			c.Set(lazy.KeyConfig, lazy.Configuration{
				DB: db.DB, Model: &model.Dog{}, Action: []lazy.Action{{Action: lazy.DefaultPostAction}}
			})
			return
		})
		g.PUT("/:id", func(c *gin.Context) {
			c.Set(lazy.KeyConfig, lazy.Configuration{
				DB: db.DB, Model: &model.Dog{}, Action: []lazy.Action{{Action: lazy.DefaultPutAction}}
			})
			return
		})
		g.PATCH("/:id", func(c *gin.Context) {
			c.Set(lazy.KeyConfig, lazy.Configuration{
				DB: db.DB, Model: &model.Dog{}, Action: []lazy.Action{{Action: lazy.DefaultPatchAction}}
			})
			return
		})
		g.DELETE("/:id", func(c *gin.Context) {
			c.Set(lazy.KeyConfig, lazy.Configuration{
				DB: db.DB, Model: &model.Dog{}, Action: []lazy.Action{{Action: lazy.DefaultDeleteAction}}
			})
			return
		})
	}
}

// CatAPI api set
func CatAPI(r *gin.Engine) {
	g := r.Group("/v1/cat").Use(gm.Trace, lazy.MiddlewareParams, lazy.MiddlewareResponse, lazy.MiddlewareDefaultResult, lazy.MiddlewareExec)
	{
		g.GET("", func(c *gin.Context) {
			c.Set(lazy.KeyConfig, lazy.Configuration{
				DB: db.DB, Model: &model.Cat{}, Action: []lazy.Action{{Action: lazy.DefaultGetAction}}
			})
			return
		})
		g.POST("", func(c *gin.Context) {
			c.Set(lazy.KeyConfig, lazy.Configuration{
				DB: db.DB, Model: &model.Cat{}, Action: []lazy.Action{{Action: lazy.DefaultPostAction}}
			})
			return
		})
		g.PUT("/:id", func(c *gin.Context) {
			c.Set(lazy.KeyConfig, lazy.Configuration{
				DB: db.DB, Model: &model.Cat{}, Action: []lazy.Action{{Action: lazy.DefaultPutAction}}
			})
			return
		})
		g.PATCH("/:id", func(c *gin.Context) {
			c.Set(lazy.KeyConfig, lazy.Configuration{
				DB: db.DB, Model: &model.Cat{}, Action: []lazy.Action{{Action: lazy.DefaultPatchAction}}
			})
			return
		})
		g.DELETE("/:id", func(c *gin.Context) {
			c.Set(lazy.KeyConfig, lazy.Configuration{
				DB: db.DB, Model: &model.Cat{}, Action: []lazy.Action{{Action: lazy.DefaultDeleteAction}}
			})
			return
		})
	}
}
