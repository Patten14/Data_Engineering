package s3bucket

import (
	"context"
	"errors"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/spf13/viper"
)

func CreateS3(s3Client *s3.Client) error {
	bucketName := viper.GetString("AWS_BUCKET_NAME")

	exists, err := BucketExists()
	if err != nil {
		log.Println(err)
		return err
	}

	if exists {
		return errors.New("bucket exists")
	}

	createBucketInput := &s3.CreateBucketInput{
		Bucket: aws.String(bucketName),
		CreateBucketConfiguration: &types.CreateBucketConfiguration{
			LocationConstraint: types.BucketLocationConstraint(viper.GetString("AWS_REGION")),
		},
	}

	_, err = s3Client.CreateBucket(context.TODO(), createBucketInput)
	if err != nil {
		log.Println(err)
		return err
	}

	log.Default().Println("Bucket created:", bucketName)
	return nil
}
