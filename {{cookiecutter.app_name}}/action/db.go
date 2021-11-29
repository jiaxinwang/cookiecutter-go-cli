package action

import (
{% if cookiecutter.use_db != "none" -%}
	"{% if cookiecutter.use_github == "y" -%}github.com/{{cookiecutter.github_username}}/{%- endif %}{{cookiecutter.app_name}}/db"
{%- endif %}
	"{% if cookiecutter.use_github == "y" -%}github.com/{{cookiecutter.github_username}}/{%- endif %}{{cookiecutter.app_name}}/model"

	"github.com/urfave/cli/v2"
)

var set = []interface{}{
	&model.Dog{},
	&model.Cat{},
}

// InitDB ...
func InitDB(c *cli.Context) error {
	Prepare(c)
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
