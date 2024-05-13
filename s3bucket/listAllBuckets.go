package s3bucket

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

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
