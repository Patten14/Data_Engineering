package s3bucket

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/spf13/viper"
)

func DeleteAllFiles(s3Client *s3.Client) {
	bucketName := viper.GetString("AWS_BUCKET_NAME")

	exists, err := BucketExists()
	if err != nil {
		log.Panic(err)
		return
	}

	if !exists {
		log.Panicf("Bucket do not exists.")
		return
	}

	objectList, err := s3Client.ListObjectsV2(context.TODO(),
		&s3.ListObjectsV2Input{Bucket: aws.String(bucketName)})

	if err != nil {
		log.Panicf("Error reading objects from bucket %v: %v", bucketName, err)
		return
	}

	count := 0
	for _, object := range objectList.Contents {
		_, err := s3Client.DeleteObject(context.TODO(), &s3.DeleteObjectInput{Bucket: aws.String(bucketName), Key: object.Key})
		if err != nil {
			log.Panicf("Error deleting object from bucket: %v", err)
			return
		}
		count++
	}
	log.Default().Printf("Deleted %v file(s)", count)
}
