package internal

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"myetc.lantron.ltd/m/internal/controller"
	"myetc.lantron.ltd/m/internal/repository/model"
	"myetc.lantron.ltd/m/internal/util"
)

type Api struct {
	Engine *gin.Engine
}

func NewApi(engine *gin.Engine) *Api {
	return &Api{
		Engine: engine,
	}
}

func (api *Api) Run(port string) {
	err := api.Engine.Run(port)
	if err != nil {
		return
	}
}

func (api *Api) Route() {
	api.configureMiddleware()

	api.Engine.GET("/ping", ping)
	api.Engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	api.Engine.GET("/api/v1/config", controller.GetConfig)

	apiGroup := api.Engine.Group("/api/v1")
	apiGroup.Use(AuthMiddleWare())
	{
		migrate(apiGroup)
		apiGroup.GET("/etc", controller.GetEtcRecordsByUserAndDate)
		apiGroup.GET("/etcSummary", controller.GetETCSummary)
	}
}

func (api *Api) configureMiddleware() {
	api.Engine.Use(
		gin.LoggerWithWriter(gin.DefaultWriter, "/ping", "/metrics"),
		gin.Recovery(),
	)
	//p := ginprometheus.NewPrometheus("myetc-api")
	//p.Use(api.Engine)
	//
	api.Engine.Use(util.Cors())
}

func ping(c *gin.Context) {
	model.ResSuccess(c, gin.H{"message": "pong"})
}

func migrate(group *gin.RouterGroup) {
	group.POST("/migration", func(context *gin.Context) {
		go func() {
			err := model.Migrate(context)
			if err != nil {
				model.ResFailureV2(context, 400, err.Error())
				return
			}
		}()
		model.ResSuccess(context, gin.H{"message": "migration successful"})
	})
}
