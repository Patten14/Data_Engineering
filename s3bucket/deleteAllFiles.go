package s3bucket

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func DeleteAllFiles(s3Client *s3.Client, bucketName string) {
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
