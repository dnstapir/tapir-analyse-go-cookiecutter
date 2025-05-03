package setup

import (
    "github.com/dnstapir/{{cookiecutter.module}}/app"
)

type AppConf struct {
	Debug      bool            `json:"debug"`
	Quiet      bool            `json:"quiet"`
}

func BuildApp(conf AppConf) (app.App, error) {
	a := app.App{}
	return a, nil
}
