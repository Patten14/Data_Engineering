package main

import (
	"DataEngineeringPipeline/s3bucket"

	"github.com/spf13/viper"
)

func InitViper() {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
}

func main() {
	InitViper()
	WorkflowS3()

}

func CreateGlue() {

}

func WorkflowS3() {
	s3Client := s3bucket.GetS3Client()

	s3bucket.CreateS3(s3Client)
	s3bucket.UploadFile(s3Client, "C:\\Users\\Patrick\\OneDrive - IU International University of Applied Sciences\\3. Semester\\Projekt Data Engineering\\daten.txt", "Testdaten")
	s3bucket.DeleteS3(s3Client, true)
}
