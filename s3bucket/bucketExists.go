package s3bucket

import (
	"context"
	"errors"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/aws/smithy-go"
	"github.com/spf13/viper"
)

func BucketExists() (bool, error) {
	s3Client := GetS3Client()
	bucketName := viper.GetString("AWS_BUCKET_NAME")

	_, err := s3Client.HeadBucket(context.TODO(), &s3.HeadBucketInput{
		Bucket: aws.String(bucketName),
	})

	exists := true

	if err != nil {
		var apiError smithy.APIError
		if errors.As(err, &apiError) {
			switch apiError.(type) {
			case *types.NotFound:
				log.Printf("Bucket %v is available.\n", bucketName)
				exists = false
				err = nil
			default:
				log.Printf("Either you don't have access to bucket %v or another error occurred. "+
					"Here's what happened: %v\n", bucketName, err)
			}
		}
	}

	return exists, err
}
