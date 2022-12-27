package main

import (
	"context"
	"log"
	"myetc.lantron.ltd/m/cmd/cron/spider"
	"myetc.lantron.ltd/m/internal/repository/model"
	"myetc.lantron.ltd/m/internal/util"
	"time"

	"github.com/robfig/cron"
)

// 0. 启动定时任务之前需要先实例一下数据库，拉取html数据之后解析好存到DB，用户访问时拿到的是本DB中的数据
// 1. 用户绑定账号之后在数据库插入一条用户数据
// 2. 拉取用户列表，每隔一段时间通过cronjob同步一次数据到数据库
// 3. 注意事项，频率不要太高，IP要动态
func main() {

	app := NewApp()
	app.Configure().ConnectMysql()

	log.Println("Starting cron job server...")
	c := cron.New()
	ctx := context.Background()
	// 5秒钟运行一次
	err := c.AddFunc("*/50 * * * * ?", func() {
		users, err := model.GetUsers(ctx)
		if err != nil {
			return
		}
		for _, user := range users {
			log.Println("start sync user: " + user.Username)
			spider.SyncUser(ctx, user.Username, user.Password)
			time.Sleep(time.Duration(util.RandomRange(10, 30)) * time.Second)
		}

	})
	if err != nil {
		return
	}
	c.Start()

	t1 := time.NewTimer(time.Second * 10)
	for {
		select {
		case <-t1.C:
			t1.Reset(time.Second * 10)
		}
	}
}
