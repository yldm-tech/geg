package internal

import (
	"context"
	"log"
	"myetc.lantron.ltd/m/internal/repository/model"
	"myetc.lantron.ltd/m/internal/repository/mysql"
	"os"

	"github.com/gin-gonic/gin"
	"myetc.lantron.ltd/m/config"
	"myetc.lantron.ltd/m/pkg/awscli"
)

type App struct{}

func NewApp() *App {
	return &App{}
}

func (app *App) Close() {
	//stream.wCloseProducer()
}

func (app *App) Configure() *App {
	env := config.Local
	if len(os.Args) > 1 {
		env = config.ENV(os.Args[1])
	}
	config.Setup(env)
	return app
}

func (app *App) ConnectDB() *App {
	if err := mysql.NewEnt(); err != nil {
		log.Fatalln("couldn't connect mysql, err: ", err)
	}
	ctx := context.Background()
	if err := model.Migrate(ctx); err != nil {
		log.Fatalln("migration failed ", err)
	}
	return app
}

func (app *App) ConfigureSDK() *App {
	if err := awscli.InitAwsSdk(); err != nil {
		log.Fatalln("couldn't initialize aws client, err: ", err)
	}
	return app
}

func (app *App) Serve() {
	r := gin.New()
	api := NewApi(r)
	api.Route()
	api.Run(config.GetServerPort())
}
