package s3bucket

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

func createS3(s3Client *s3.Client, bucketName string) error {
	_, err := s3Client.CreateBucket(context.TODO(),
		&s3.CreateBucketInput{Bucket: aws.String(bucketName),
			CreateBucketConfiguration: &types.CreateBucketConfiguration{LocationConstraint: types.BucketLocationConstraint(region)},
		})
	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println("Created:", bucketName)
	return nil
}
