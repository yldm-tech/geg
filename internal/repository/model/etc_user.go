package model

import (
	"context"
	"gopkg.in/guregu/null.v4"
	"myetc.lantron.ltd/m/ent"
	"myetc.lantron.ltd/m/internal/repository/domain"
	"myetc.lantron.ltd/m/internal/repository/mysql"
)

type ETCUser struct {
	Id       int64       `json:"id"`
	Username string      `json:"username"`
	Password string      `json:"password"`
	Icon     null.String ` json:"icon"`
}

// ToDomain
// do not use pointer
func (g ETCUser) ToDomain() *domain.ETCUser {
	var user = domain.ETCUser{
		Id:       g.Id,
		Username: g.Username,
	}
	return &user
}

func GetUsers(ctx context.Context) ([]*ent.ETCUser, error) {
	users, err := mysql.Client().ETCUser.Query().All(ctx)
	if err != nil {
		return nil, err
	}
	return users, nil
}
