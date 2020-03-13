package action

import (
	"github.com/gin-gonic/gin"
	"github.com/urfave/cli/v2"
)

var Server = cli.Command{
	Name:  "server",
	Usage: "http server",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "conf, c",
			Value: "config.toml",
		},
	},
	Action: run,
}

func run(c *cli.Context) {
	GinEngine().Run("")
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
