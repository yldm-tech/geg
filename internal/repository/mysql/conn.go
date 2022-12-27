package mysql

import (
	"entgo.io/ent/dialect"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
	"myetc.lantron.ltd/m/config"
	"myetc.lantron.ltd/m/ent"
)

var (
	client *ent.Client
)

func NewEnt() error {
	connectCfg := config.GetConfig().Database
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true",
		connectCfg.User, connectCfg.Password, connectCfg.Host, connectCfg.Port, connectCfg.DB,
	)
	var err error
	var opts []ent.Option
	if config.GetConfig().Debug {
		opts = append(opts, ent.Debug())
	}
	client, err = OpenEnt(dialect.MySQL, dsn, opts...)
	if err != nil {
		logrus.Error("init db error...")
	}
	//defer client.Close()
	return err
}

func Client() *ent.Client {
	return client
}

func MustNoErr(err error) {
	if err != nil {
		logrus.Errorln("unexpected error occurred ", err)
		panic(err)
	}
}
func OpenEnt(driver string, dsn string, opts ...ent.Option) (*ent.Client, error) {
	return ent.Open(driver, dsn, opts...)
}
