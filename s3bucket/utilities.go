package s3bucket

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/spf13/viper"
)

func GetS3Client() *s3.Client {
	sdkConfig, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(viper.GetString("AWS_REGION")))
	if err != nil {
		log.Panic("Couldn't load default configuration. Have you set up your AWS account?")
		log.Panic(err)
		panic(err)
	}
	return s3.NewFromConfig(sdkConfig)
}
