/**
 * @Author: Evan<suzukaze.hazuki2020@gmail.com>
 * @Description: app的相关设置，包括启动端口，日志存储路径，运行环境等
 * @File:  app
 * @Version: 1.0.0
 * @Date: 2022/11/30 14:50
 */

package config

type ENV string

const (
	Local ENV = "local"
	Stg   ENV = "stg"
	Prod  ENV = "prod"
)

func (e ENV) Str() string {
	return string(e)
}

type Config struct {
	App      App
	Database Database
	Debug    bool
}

type App struct {
	ApiPort     string
	LogFileName string
	S3Bucket    string
	AwsProfile  string
	Env         ENV
	Domain      string
}

var config *Config

func GetConfig() *Config {
	return config
}

type Database struct {
	Type     string
	User     string
	Password string
	Host     string
	Port     int
	DB       string
}

func GetServerPort() string {
	return ":" + config.App.ApiPort
}
