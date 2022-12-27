package awscli

import (
	"context"
	"sync"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"myetc.lantron.ltd/m/config"
)

var (
	s3Client *s3.Client
	once     sync.Once
)

func InitAwsSdk() error {
	var (
		cfg aws.Config
		err error
	)
	appConfig := config.GetConfig().App
	profile, region := appConfig.AwsProfile, "ap-northeast-1"

	if appConfig.Env == config.Local {
		cfg, err = awsConfig.LoadDefaultConfig(context.TODO(), awsConfig.WithSharedConfigProfile(profile), awsConfig.WithRegion(region))
	} else {
		cfg, err = awsConfig.LoadDefaultConfig(context.TODO(), awsConfig.WithRegion(region))
	}
	if err != nil {
		return err
	}
	once.Do(func() {
		s3Client = s3.NewFromConfig(cfg)
	})
	return nil
}
