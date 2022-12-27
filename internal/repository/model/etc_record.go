package model

import (
	"context"
	"myetc.lantron.ltd/m/ent"
	"myetc.lantron.ltd/m/ent/etcrecord"
	"myetc.lantron.ltd/m/internal/constant"
	"myetc.lantron.ltd/m/internal/repository/domain"
	"myetc.lantron.ltd/m/internal/repository/mysql"
	"myetc.lantron.ltd/m/internal/util"
)

type ETCRecord struct {
	Username      string  `json:"username"`
	Entry         string  `json:"entry,omitempty"`
	EntryYear     int32   `json:"entryYear,omitempty"`
	EntryMonth    int32   `json:"entryMonth,omitempty"`
	EntryDay      int32   `json:"entryDay,omitempty"`
	EntryTime     string  `json:"entryTime,omitempty"`
	Exit          string  `json:"exit"`
	ExitDate      string  `json:"exitDate"`
	ExitTime      string  `json:"exitTime"`
	TotalPrice    int32   `json:"totalPrice"`
	DiscountPrice int32   `json:"discountPrice"`
	PaidPrice     int32   `json:"paidPrice"`
	CarType       int8    `json:"carType"`
	CarNumber     string  `json:"carNumber"`
	CardNumber    string  `json:"cardNumber"`
	Status        string  `json:"status"`
	Comment       *string `json:"comment"`
}

func BatchInsertEtc(ctx context.Context, etcList []*ETCRecord) error {
	if etcList == nil || len(etcList) == 0 {
		return nil
	}
	bulk := make([]*ent.ETCRecordCreate, len(etcList))
	for i, etc := range etcList {
		bulk[i] = mysql.Client().ETCRecord.Create().
			SetUsername(etc.Username).
			SetEntry(etc.Entry).
			SetEntryYear(etc.EntryYear).
			SetEntryMonth(etc.EntryMonth).
			SetEntryDay(etc.EntryDay).
			SetExit(etc.Exit).
			SetExitDate(etc.ExitDate).
			SetExitTime(etc.ExitTime).
			SetTotalPrice(etc.TotalPrice).
			SetDiscountPrice(etc.DiscountPrice).
			SetPaidPrice(etc.PaidPrice).
			SetCarType(etc.CarType).
			SetCarNumber(etc.CardNumber).
			SetCardNumber(etc.CardNumber).
			SetStatus(etc.Status).
			SetNillableComment(etc.Comment)

	}
	err := mysql.Client().ETCRecord.CreateBulk(bulk...).Exec(ctx)
	return err
}

func GetEtcRecordsByUserAndDate(ctx context.Context,
	username string,
	year int32,
	month int32,
) ([]*ent.ETCRecord, error) {
	query := mysql.Client().ETCRecord.Query().
		Where(
			etcrecord.Username(username),
			etcrecord.EntryYear(year))
	if month > 0 {
		query.Where(etcrecord.EntryMonth(month))
	}
	records, err := query.
		All(ctx)
	if err != nil {
		return nil, err
	}
	return records, nil
}

func GetETCSummary(
	ctx context.Context,
	username string,
	summaryType string,
) (*domain.ETCSummary, error) {
	var record *domain.ETCSummary
	db := mysql.Client().ETCRecord.Query().Where(
		etcrecord.Username(username),
	)

	if summaryType == constant.QueryByYear {
		db.Where(etcrecord.EntryYear(int32(util.GetYear())))
	}

	if summaryType == constant.QueryByMonth {
		db.Where(
			etcrecord.EntryYear(int32(util.GetYear())),
			etcrecord.EntryMonth(int32(11))) //fixme 测试用，当月没有数据
	}

	// 总个数
	totalCount, err := db.Count(ctx)
	if err != nil {
		totalCount = 0
	}

	// 总金额
	totalAmount, err := db.Aggregate(ent.Sum(etcrecord.FieldPaidPrice)).Int(ctx)
	if err != nil {
		totalAmount = 0
	}

	// 条形图
	var v []domain.Strip
	group := etcrecord.FieldEntryMonth
	if summaryType == constant.QueryByAll {
		group = etcrecord.FieldEntryYear
	} else if summaryType == constant.QueryByMonth {
		group = etcrecord.FieldEntryDay
	}
	err = db.
		GroupBy(group).
		Aggregate(ent.Sum(etcrecord.FieldPaidPrice)).Scan(ctx, &v)
	if err != nil {
		return nil, err
	}

	// 饼状图
	var pipeEntry []domain.Pie

	err = db.GroupBy(etcrecord.FieldEntry).
		Aggregate(ent.Sum(etcrecord.FieldPaidPrice)).
		Scan(ctx, &pipeEntry)
	if err != nil {
		return nil, err
	}

	var pipeExit []domain.Pie
	err = db.GroupBy(etcrecord.FieldExit).
		Aggregate(ent.Sum(etcrecord.FieldPaidPrice)).
		Scan(ctx, &pipeExit)
	if err != nil {
		return nil, err
	}

	// 汇总
	record = &domain.ETCSummary{
		Amount:    totalAmount,
		Count:     totalCount,
		Strip:     v,
		PieEntry:  pipeEntry,
		PieExit:   pipeExit,
		QueryType: summaryType,
	}
	return record, nil

}
