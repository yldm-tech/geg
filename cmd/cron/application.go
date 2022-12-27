package main

import (
	log "github.com/sirupsen/logrus"
	"myetc.lantron.ltd/m/config"
	"myetc.lantron.ltd/m/internal/repository/mysql"
	"os"
)

type App struct{}

func NewApp() *App {
	return &App{}
}

func (app *App) Configure() *App {
	env := config.Local
	if len(os.Args) > 1 {
		env = config.ENV(os.Args[1])
	}
	config.Setup(env)
	return app
}

func (app *App) ConnectMysql() *App {
	if err := mysql.Connect(); err != nil {
		log.Fatalln("couldn't connect mysql, err: ", err)
		return app
	}
	return app
}
