package action

import (
	"{% if cookiecutter.use_github == "y" -%}github.com/{{cookiecutter.github_username}}/{%- endif %}{{cookiecutter.app_name}}/db"
	"{% if cookiecutter.use_github == "y" -%}github.com/{{cookiecutter.github_username}}/{%- endif %}{{cookiecutter.app_name}}/model"
	"{% if cookiecutter.use_github == "y" -%}github.com/{{cookiecutter.github_username}}/{%- endif %}{{cookiecutter.app_name}}/config"

	"github.com/urfave/cli/v2"
)

func Prepare(c *cli.Context) {
	config.Load(c.String(`conf`))
	db.Connect(config.Config.Database.DSN)
	model.SetDB(db.DB)
}
