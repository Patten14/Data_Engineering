package glue

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/glue"
	"github.com/spf13/viper"
)

func GetClient() *glue.Client {
	sdkConfig, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(viper.GetString("AWS_REGION")))
	if err != nil {
		log.Panic(err)
	}

	client := glue.NewFromConfig(sdkConfig)
	return client
}
