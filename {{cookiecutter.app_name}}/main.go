package main

import (
	"fmt"
	"io"
	"os"
	"{% if cookiecutter.use_github == "y" -%}github.com/{{cookiecutter.github_username}}/{%- endif %}{{cookiecutter.app_name}}/action"
	"{% if cookiecutter.use_github == "y" -%}github.com/{{cookiecutter.github_username}}/{%- endif %}{{cookiecutter.app_name}}/util/l"
{% if cookiecutter.use_survey == "y" -%}
	"github.com/AlecAivazis/survey/v2"
{%- endif %}
	"github.com/urfave/cli/v2"
)

{% if cookiecutter.use_survey == "y" -%}
var questions = []*survey.Question{
	{
		Name: "here",
		Prompt: &survey.Input{
			Message: "Am I sitting in a tin can",
		},
		Validate:  survey.Required,
		Transform: survey.Title,
	},
	{
		Name: "color",
		Prompt: &survey.Select{
			Message: "planet earth is:",
			Options: []string{"red", "blue", "green"},
		},
		Validate: survey.Required,
	},
}
{%- endif %}

func main() {
{% if cookiecutter.use_gin == "y" -%}
	gin.DefaultWriter = mw
{%- endif %}

	app := cli.NewApp()
	app.Name = "{{cookiecutter.app_name}}"
	app.Version = "0.1.0"
	app.Commands = []*cli.Command{
		&echo,
{% if cookiecutter.use_gin == "y" -%}		
		&server,
{%- endif %}
{% if cookiecutter.use_db != "none" -%}		
		&database,
{%- endif %}
	}
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:    `star`,
			Aliases: []string{`s`},
			Value:   "the stars look very different today",
		},
		&cli.StringFlag{
			Name:  "conf",
			Aliases: []string{`c`},
			Value: "./config.toml",
		},
{% if cookiecutter.use_db != "none" -%}
		&cli.StringFlag{
			Name:     `dsn`,
			Aliases:  []string{`d`},
			Usage:    "data source name",
			Value:    "",
			FilePath: "./dsn",
		},
{%- endif %}
	}

{% if cookiecutter.use_survey == "y" -%}
	answers := struct {
		Name  string
		Color string
	}{}
	if err := survey.Ask(questions, &answers);err != nil {
		l.S.Panic(err)
		return
	}
{%- endif %}

	err = app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}

var echo = cli.Command{
	Name:  "echo",
	Usage: "3... 2... 1... liftoff",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name: "d",
		},
	},
	Action: action.Echo,
}

{% if cookiecutter.use_gin == "y" -%}
var server = cli.Command{
	Name:  "server",
	Usage: "http server",
	Flags: []cli.Flag{},
	Action: action.Server,
}
{%- endif %}

{% if cookiecutter.use_db != "none" -%}		
var database = cli.Command{
	Name:  "db",
	Usage: "db",
	Flags: []cli.Flag{},
	Action: action.InitDB,
}
{%- endif %}
