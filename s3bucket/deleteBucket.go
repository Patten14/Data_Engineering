package s3bucket

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/spf13/viper"
)

func DeleteS3(s3Client *s3.Client, forceDelete bool) error {
	bucketName := viper.GetString("AWS_BUCKET_NAME")

	exists, err := BucketExists()
	if err != nil {
		log.Panic(err)
		return err
	}

	if !exists {
		log.Panic("Bucket do not exists.")
		return err
	}

	if forceDelete {
		DeleteAllFiles(s3Client)
	}

	_, err = s3Client.DeleteBucket(context.TODO(),
		&s3.DeleteBucketInput{Bucket: aws.String(bucketName)})
	if err != nil {
		log.Panic(err)
		return err
	}
	log.Println("Deleted:", bucketName)
	return nil
}
