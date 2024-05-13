package s3bucket

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func deleteS3(s3Client *s3.Client, bucketName string) error {
	_, err := s3Client.DeleteBucket(context.TODO(),
		&s3.DeleteBucketInput{Bucket: aws.String(bucketName)})
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("Deleted:", bucketName)
	return nil
}
