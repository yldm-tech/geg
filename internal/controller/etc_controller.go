package controller

import (
	"github.com/gin-gonic/gin"
	"myetc.lantron.ltd/m/internal/constant"
	"myetc.lantron.ltd/m/internal/repository/model"
	"myetc.lantron.ltd/m/internal/service"
	"myetc.lantron.ltd/m/internal/util"
	"strconv"
)

func GetEtcRecordsByUserAndDate(c *gin.Context) {
	username, yearStr, monthStr := c.Query("username"),
		c.DefaultQuery("year", util.GetYearString()),
		c.Query("month")

	year, err := strconv.Atoi(yearStr)
	if err != nil {
		model.ResFailureV2(c, 400, err.Error())
		return
	}
	month, err := strconv.Atoi(monthStr)
	if err != nil {
		month = 0
	}

	records, err := service.GetETC(c, username, int32(year), int32(month))
	if err != nil {
		model.ResFailureV2(c, 400, err.Error())
		return
	}

	model.ResSuccess(c, records)
}

func GetETCSummary(c *gin.Context) {
	username, summaryType := c.Query("username"),
		c.DefaultQuery("queryType", constant.QueryByYear)

	records, err := service.GetETCSummary(c, username, summaryType)
	if err != nil {
		model.ResFailureV2(c, 400, err.Error())
		return
	}

	model.ResSuccess(c, records)
}
