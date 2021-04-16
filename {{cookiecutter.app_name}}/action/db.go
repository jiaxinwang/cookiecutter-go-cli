package action

import (
	"{% if cookiecutter.use_github == "y" -%}github.com/{{cookiecutter.github_username}}/{%- endif %}{{cookiecutter.app_name}}/db"

	"github.com/urfave/cli/v2"
)

var set = []interface{}{}

// InitDB ...
func InitDB(c *cli.Context) error {
	db.Connect(c.String(`dsn`))
	if c.Bool("recreate") {
		for _, v := range set {
			db.DB.Migrator().DropTable(v)
		}
	}
	for _, v := range set {
		db.DB.Migrator().AutoMigrate(v)
	}
	return nil
}
coo