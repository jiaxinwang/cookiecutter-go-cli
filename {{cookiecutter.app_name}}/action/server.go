package action

import (
	"github.com/gin-gonic/gin"
	"github.com/urfave/cli/v2"
)

// Server ...
func Server(c *cli.Context) error {
	GinEngine().Run("")
	return nil
}

// GinEngine ...
func GinEngine() *gin.Engine {
	var r *gin.Engine
	r = gin.Default()
	r.GET("/health")
	V1(r)
	return r
}

// V1 api set
func V1(r *gin.Engine) {
	g := r.Group("/v1")
	{
		g.GET("/echo", nil)
	}

}
