package main

import (
	"DataEngineeringPipeline/glue"
	"DataEngineeringPipeline/s3bucket"

	awsglue "github.com/aws/aws-sdk-go-v2/service/glue"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/spf13/viper"
)

func InitViper() {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
}

func main() {
	s3Client := s3bucket.GetClient()
	glueClient := glue.GetClient()

	InitViper()
	WorkflowS3(s3Client)
	CreateGlue(glueClient)
}

func CreateGlue(glueClient *awsglue.Client) {

	err := glue.CreateDatabase(glueClient)
	if err != nil {
		return
	}
	err = glue.CreateCrawler(glueClient)
	if err != nil {
		return
	}
	err = glue.CreateJob(glueClient)
	if err != nil {
		return
	}
	err = glue.CreateTrigger(glueClient)
	if err != nil {
		return
	}
}

func WorkflowS3(s3Client *s3.Client) {
	s3bucket.CreateS3(s3Client)
	s3bucket.UploadFile(s3Client, "C:\\Users\\Patrick\\OneDrive - IU International University of Applied Sciences\\3. Semester\\Projekt Data Engineering\\data.csv", "data/Testdaten.csv")
	s3bucket.UploadFile(s3Client, "D:\\Documents\\GitHub\\Data_Engineering\\Scripts\\glueJob.py", "scripts/glueJob.py")
}
