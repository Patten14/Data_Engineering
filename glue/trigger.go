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

func CreateTrigger(glueClient *glue.Client) error {
	createTriggerOutput, err := glueClient.CreateTrigger(context.TODO(), &glue.CreateTriggerInput{
		Name: aws.String("GlueJobTrigger"),
		Type: "CONDITIONAL",
		Actions: []types.Action{
			{
				JobName: aws.String(viper.GetString("AWS_GLUE_JOB_NAME")),
			},
		},
		Predicate: &types.Predicate{
			Logical: types.LogicalAny,
			Conditions: []types.Condition{
				{
					CrawlState:      types.CrawlStateSucceeded,
					CrawlerName:     aws.String(viper.GetString("AWS_GLUE_CRAWLER_NAME")),
					LogicalOperator: types.LogicalOperatorEquals,
				},
			},
		},
		StartOnCreation: true,
	})

	if err != nil {
		var alreadyExistsError *types.AlreadyExistsException
		if !errors.As(err, &alreadyExistsError) {
			log.Println("Failed to create Glue job trigger.", err)
			return err
		}
	}
	log.Println("Create Glue job trigger success", createTriggerOutput.Name)
	return nil
}
