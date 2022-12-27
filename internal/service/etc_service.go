package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"myetc.lantron.ltd/m/internal/repository/domain"
	"myetc.lantron.ltd/m/internal/repository/model"
	"myetc.lantron.ltd/m/internal/util"
	"strings"
)

func GetETC(c *gin.Context, username string, year int32, month int32) (map[string][]interface{}, error) {
	if len(strings.Trim(username, " ")) == 0 {
		model.ResFailureV2(c, 400, "please input username")
		return nil, errors.New("please input username")
	}

	records, err := model.GetEtcRecordsByUserAndDate(c, username, year, month)
	if err != nil {
		return nil, err
	}

	etcMap := util.ListToMap(records, "ExitDate")

	return etcMap, nil
}

func GetETCSummary(c *gin.Context, username, summaryType string) (*domain.ETCSummary, error) {
	if len(strings.Trim(username, " ")) == 0 {
		model.ResFailureV2(c, 400, "please input username")
		return nil, errors.New("please input username")
	}

	records, err := model.GetETCSummary(c, username, summaryType)
	if err != nil {
		return nil, err
	}

	return records, nil
}
