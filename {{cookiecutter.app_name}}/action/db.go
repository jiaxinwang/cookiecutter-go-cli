package action

import (
	"{% if cookiecutter.use_github == "y" -%}github.com/{{cookiecutter.github_username}}/{%- endif %}{{cookiecutter.app_name}}/db"

	"github.com/urfave/cli/v2"
)

func initDB(c *cli.Context) error {
	db.Connect(c.String(`dsn`))
	if c.Bool("recreate") {
		db.DB.DropTable()
	}
	db.DB.AutoMigrate()

	return nil
}
