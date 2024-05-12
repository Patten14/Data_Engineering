package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

var region string = "us-west-2"
var bucketName string = "tf-dataengineering-patrick"

func main() {
	s3Client, err := GetS3Client()
	if err != nil {
		return
	}

	createS3(s3Client, bucketName)
	UploadFile(s3Client, bucketName, "C:\\Users\\Patrick\\OneDrive - IU International University of Applied Sciences\\3. Semester\\Projekt Data Engineering\\daten.txt", "Testdaten")
	deleteS3(s3Client, bucketName)

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

func ListAllBuckets(s3Client *s3.Client) {
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

func UploadFile(s3Client *s3.Client, bucketName string, filePath string, fileName string) error {
	file, err := os.Open(filePath)
	if err != nil {
		log.Printf("Couldn't open file %v to upload. Here's why: %v\n", filePath, err)
	} else {
		defer file.Close()
		_, err = s3Client.PutObject(context.TODO(), &s3.PutObjectInput{
			Bucket: aws.String(bucketName),
			Key:    aws.String(fileName),
			Body:   file,
		})
		if err != nil {
			log.Printf("Couldn't upload file %v to %v:%v. Here's why: %v\n",
				fileName, bucketName, fileName, err)
		}
	}
	return err
}

func ListAllFiles(s3Client *s3.Client, bucketName string) {
	objectList, err := s3Client.ListObjectsV2(context.TODO(),
		&s3.ListObjectsV2Input{Bucket: aws.String(bucketName)})

	if err != nil {
		fmt.Printf("Error reading objects from bucket %v: %v\n", bucketName, err)
		return
	}

	for _, object := range objectList.Contents {
		outputList, err := s3Client.DeleteObject(context.TODO(), &s3.DeleteObjectInput{Bucket: aws.String(bucketName), Key: object.Key})
		if err != nil {
			fmt.Printf("Error deleting object from bucket: %v\n", err)
			return
		}
		fmt.Println("File", object.Key, "is deleted", outputList.DeleteMarker)
	}
}
