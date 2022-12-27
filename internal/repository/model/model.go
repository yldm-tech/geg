package model

import (
	"context"
	log "github.com/sirupsen/logrus"
	"myetc.lantron.ltd/m/ent/migrate"
	"myetc.lantron.ltd/m/internal/repository/mysql"
)

func Migrate(ctx context.Context) error {
	err := mysql.Client().Schema.Create(ctx,
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	)
	if err != nil {
		log.Fatalf("fialed create schema resource:%v", err)
		return err
	}
	return nil
}
