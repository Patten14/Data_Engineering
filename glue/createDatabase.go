package glue

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/glue"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	"github.com/spf13/viper"
)

func CreateDatabase() error {
	glueClient := GetClient()

	databaseName := viper.GetString("AWS_GLUE_DATABASE_NAME")

	createDatabaseInput := &glue.CreateDatabaseInput{
		DatabaseInput: &types.DatabaseInput{
			Name: aws.String(databaseName),
		},
	}

	_, err := glueClient.CreateDatabase(context.TODO(), createDatabaseInput)
	if err != nil {
		log.Panic("Failed to create Glue database")
		log.Panic(err)
		return err
	}
	log.Println("Created Glue database:", databaseName)
	return nil
}
