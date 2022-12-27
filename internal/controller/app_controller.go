package controller

import (
	"github.com/gin-gonic/gin"
	"myetc.lantron.ltd/m/config"
	"myetc.lantron.ltd/m/internal/repository/model"
)

// GetConfig
// @Tags App管理
// @Summary 获取客户端所需配置(v1.2.0)
// @Accept  json
// @Produce json
// @Router /api/v1/config [get]
func GetConfig(c *gin.Context) {
	app := config.GetConfig().App
	type ClientConfig struct {
		Env string `json:"env"`
	}
	model.ResSuccess(c, ClientConfig{
		Env: app.Env.Str(),
	})
}
