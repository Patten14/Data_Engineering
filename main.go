package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

var region string = "us-west-2"
var bucketName string = "tf-dataengineering-patrick"

func main() {
	createS3(bucketName)
	ListAllBuckets()
	// deleteS3(bucketName)
	// ListAllBuckets()

}

func GetS3Client() (*s3.Client, error) {
	sdkConfig, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
	if err != nil {
		fmt.Println("Couldn't load default configuration. Have you set up your AWS account?")
		fmt.Println(err)
		return &s3.Client{}, err
	}
	return s3.NewFromConfig(sdkConfig), err
}

func deleteS3(bucketName string) error {
	s3Client, err := GetS3Client()
	if err != nil {
		return err
	}

	_, err = s3Client.DeleteBucket(context.TODO(),
		&s3.DeleteBucketInput{Bucket: aws.String(bucketName)})
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("Deleted:", bucketName)
	return nil
}

func createS3(bucketName string) error {
	s3Client, err := GetS3Client()
	if err != nil {
		return err
	}

	_, err = s3Client.CreateBucket(context.TODO(),
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

func ListAllBuckets() {
	s3Client, err := GetS3Client()
	if err != nil {
		return
	}

	fmt.Println("Let's list up all buckets for your account.")
	result, err := s3Client.ListBuckets(context.TODO(), &s3.ListBucketsInput{})
	if err != nil {
		fmt.Printf("Couldn't list buckets for your account. Here's why: %v\n", err)
		return
	}
	if len(result.Buckets) == 0 {
		fmt.Println("You don't have any buckets!")
	} else {
		for _, bucket := range result.Buckets[:len(result.Buckets)] {
			fmt.Printf("\t%v\n", *bucket.Name)
		}
	}
}

func UploadFile(bucketName string, input string) error {

	return nil
}
