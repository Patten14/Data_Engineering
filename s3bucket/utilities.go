package s3bucket

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

var region string = "us-west-2"
var bucketName string = "tf-dataengineering-patrick"

func GetS3Client() *s3.Client {
	sdkConfig, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
	if err != nil {
		fmt.Println("Couldn't load default configuration. Have you set up your AWS account?")
		fmt.Println(err)
		panic(err)
	}
	return s3.NewFromConfig(sdkConfig)
}
