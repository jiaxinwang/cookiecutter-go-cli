package main

import (
	"fmt"
	"io"
	"os"
	"{% if cookiecutter.use_github == "y" -%}github.com/{{cookiecutter.github_username}}/{%- endif %}{{cookiecutter.app_name}}/action"

	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

func main() {
	logrus.SetFormatter(&nested.Formatter{
		TimestampFormat: "01/02/15:04:05",
		HideKeys:        false,
		ShowFullLevel:   true,
		FieldsOrder:     []string{"component", "category"},
	})
	logrus.SetReportCaller(true)
	logrus.SetLevel(logrus.TraceLevel)
	logFile, err := os.OpenFile("logs.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(fmt.Sprintf("[Error]: %s", err))
	}
	mw := io.MultiWriter(os.Stdout, logFile)
	logrus.SetOutput(mw)

	app := cli.NewApp()
	app.Name = "{{cookiecutter.app_name}}"
	app.Version = "0.1.0"
	app.Commands = []*cli.Command{
		&action.Echo,
	}
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:  "star, s",
			Value: "the stars look very different today",
		},
	}

	err = app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}
