package model

import (
	"context"
	"myetc.lantron.ltd/m/ent"
	"myetc.lantron.ltd/m/ent/etcsyncrecord"
	"myetc.lantron.ltd/m/internal/repository/mysql"
)

type ETCSyncRecord struct {
	Username string `json:"username"`
	SyncDate string `json:"sync_data"` // yyyyMM(202211)
}

func AddEtcSyncRecord(ctx context.Context, record ETCSyncRecord) error {
	_, err := mysql.Client().ETCSyncRecord.Create().
		SetUsername(record.Username).
		SetSyncData(record.SyncDate).
		Save(ctx)
	if err != nil {
		return err
	}
	return nil
}

func GetEtcSyncRecord(ctx context.Context, username string, date string) (*ent.ETCSyncRecord, error) {
	record, err := mysql.Client().ETCSyncRecord.Query().Where(
		etcsyncrecord.Username(username),
		etcsyncrecord.SyncData(date),
	).First(ctx)
	if err != nil {
		return nil, err
	}
	return record, nil
}
