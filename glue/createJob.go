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

func CreateJob(glueClient *glue.Client) error {
	jobName := viper.GetString("AWS_GLUE_JOB_NAME")
	pythonScriptPath := (viper.GetString("AWS_BUCKET_PATH") + "/scripts/glueJob.py")
	pythonVersion := "3"
	role := viper.GetString("AWS_GLUE_CRAWLER_ROLE_ARN")

	command := &types.JobCommand{
		Name:           aws.String("glueetl"),
		ScriptLocation: aws.String(pythonScriptPath),
		PythonVersion:  aws.String(pythonVersion),
	}

	createJobInput := &glue.CreateJobInput{
		Name:            aws.String(jobName),
		Role:            aws.String(role),
		GlueVersion:     aws.String("4.0"),
		WorkerType:      types.WorkerTypeG1x,
		NumberOfWorkers: aws.Int32(2),
		Command:         command,
	}

	createJobOutput, err := glueClient.CreateJob(context.TODO(), createJobInput)
	if err != nil {
		var alreadyExistsError *types.AlreadyExistsException
		if !errors.As(err, &alreadyExistsError) {
			log.Println("Failed to create Glue job", err)
			return err
		}
	}

	log.Println("Create Glue job success", createJobOutput.Name)
	return nil
}
