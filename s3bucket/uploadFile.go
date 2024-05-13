package s3bucket

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/spf13/viper"
)

func UploadFile(s3Client *s3.Client, filePath string, fileName string) error {
	bucketName := viper.GetString("AWS_BUCKET_NAME")
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
