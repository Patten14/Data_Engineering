package glue

import (
	"context"
	"errors"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/glue"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	"github.com/spf13/viper"
)

func CreateCrawler(glueClient *glue.Client) error {

	crawlerName := viper.GetString("AWS_GLUE_CRAWLER_NAME")
	roleArn := viper.GetString("AWS_GLUE_CRAWLER_ROLE_ARN")
	bucketPath := viper.GetString("AWS_BUCKET_NAME") + "/data/"
	dbName := viper.GetString("AWS_GLUE_DATABASE_NAME")
	schedule := viper.GetString("AWS_GLUE_CRAWLER_SCHEDULE")

	createCrawlerInput := &glue.CreateCrawlerInput{
		Name:         aws.String(crawlerName),
		Role:         aws.String(roleArn),
		DatabaseName: aws.String(dbName),
		Schedule:     aws.String(schedule),
		Targets: &types.CrawlerTargets{
			S3Targets: []types.S3Target{
				{
					Path: aws.String(bucketPath),
				},
			},
		},
	}

	_, err := glueClient.CreateCrawler(context.TODO(), createCrawlerInput)
	if err != nil {
		var alreadyExistsError *types.AlreadyExistsException
		if !errors.As(err, &alreadyExistsError) {
			log.Println("Failed to create Glue crawler.", err)
			return err
		}
	}
	log.Println("Glue crawler created.")

	startCrawlerInput := &glue.StartCrawlerInput{
		Name: aws.String(crawlerName),
	}
	_, err = glueClient.StartCrawler(context.TODO(), startCrawlerInput)
	if err != nil {
		log.Println("Start crawler failed.", err)
		return err
	}

	return nil
}
