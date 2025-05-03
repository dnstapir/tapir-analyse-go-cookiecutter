package app

import (
	"errors"
    "github.com/dnstapir/{{cookiecutter.module}}/app/ext"
)

type App struct {
	Log       ext.Logger
	isRunning bool
}

func (a *App) Run() <-chan error {
	doneChan := make(chan error, 10)
	a.isRunning = true

	if a.Log == nil {
		doneChan <- errors.New("no logger object")
		a.isRunning = false
	}

	return doneChan
}

func (a *App) Stop() error {
	if !a.isRunning {
		return errors.New("called stop while app was not running")
	}

	return nil
}
