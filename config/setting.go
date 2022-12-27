/**
 * @Author: Evan<suzukaze.hazuki2020@gmail.com>
 * @Description: 将配置文件映射到实体实
 * @File:  setting
 * @Version: 1.0.0
 * @Date: 2022/11/30 14:50
 */

package config

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func Setup(env ENV) {
	if err := loadConfig(env); err != nil {
		log.Fatalf("setting.Setup, fail to parse config file: %v", err)
	}
	configureLog(env)
}

func loadConfig(env ENV) error {
	viper.SetConfigName("config_" + env.Str())
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./local")
	viper.AddConfigPath("./config")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	if err := viper.Unmarshal(&config); err != nil {
		return err
	}
	config.App.Env = env
	return nil
}

func configureLog(env ENV) {
	// 最长保存1年时间，每天分割新的日志文件 yyyyMMdd

	log.SetReportCaller(false) //将函数名和行数放在日志里面
	log.SetLevel(log.InfoLevel)
	if env.Str() != "local" {
		log.SetFormatter(&log.JSONFormatter{})
	}
	//  暂时不需要写入到文件
	//path := config.App.LogFileName
	//writer, err := rotatelogs.New(
	//	path+".%Y%m%d",
	//	rotatelogs.WithLinkName(path),             // 生成软链，指向最新日志文件
	//	rotatelogs.WithMaxAge(time.Hour*24*3),     // 文件最大保存时间
	//	rotatelogs.WithRotationTime(time.Hour*24), // 日志切割时间间隔
	//)
	//if err != nil {
	//	log.Fatal("Init failed, err:", err)
	//}
	//log.SetOutput(writer)
}
