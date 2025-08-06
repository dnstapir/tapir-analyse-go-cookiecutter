package setup

import (
	"github.com/dnstapir/{{cookiecutter.module}}/app"
	"github.com/dnstapir/{{cookiecutter.module}}/internal/logging"
	"github.com/dnstapir/{{cookiecutter.module}}/internal/nats"
)

type AppConf struct {
	Debug bool       `toml:"debug"`
	Quiet bool       `toml:"quiet"`
	Nats  NatsConfig `toml:"nats"`
}

type NatsConfig struct {
	Url        string `toml:"url"`
	InSubject  string `toml:"in_subject"`
	OutSubject string `toml:"out_subject"`
	Queue      string `toml:"queue"`
}

func BuildApp(conf AppConf) (*app.App, error) {
	log := logging.Create(conf.Debug, conf.Quiet)

	natsClient, err := nats.CreateClient(conf.Nats.Url, conf.Nats.InSubject, conf.Nats.OutSubject, conf.Nats.Queue)
	if err != nil {
		return nil, err
	}

	a := app.App{
		Log:  log,
		Nats: natsClient,
	}

	return &a, nil
}
