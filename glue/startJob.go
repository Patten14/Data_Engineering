package glue

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/glue"
	"github.com/spf13/viper"
)

func StartJob(glueClient *glue.Client) {
	jobName := viper.GetString("AWS_GLUE_JOB_NAME")

	input := &glue.StartJobRunInput{
		JobName: aws.String(jobName),
	}

	_, err := glueClient.StartJobRun(context.TODO(), input)
	if err != nil {
		log.Println("Failed to start Glue job", err)
		return
	}

	log.Println("Glue job started.")
}
