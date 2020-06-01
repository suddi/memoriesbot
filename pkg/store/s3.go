package store

import (
	"bytes"
	"fmt"
	"memoriesbot/pkg/config"
	"memoriesbot/pkg/logger"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func Download(filename) ([]byte, error) {
	c := config.Get()

	service := session.Must(session.NewSession(&aws.Config{
		Region: &c.Aws.Region,
	}))
	downloader := s3manager.NewDownloader(service)
	downloader.Download(f, &s3.GetObjectInput{
		Bucket: aws.String(myBucket),
		Key:    aws.String(myString),
	})
}

// Upload - uploads file to S3
func Upload(filename string, content []byte) bool {
	c := config.Get()
	body := bytes.NewReader(content)

	service := session.Must(session.NewSession(&aws.Config{
		Region: &c.Aws.Region,
	}))

	uploader := s3manager.NewUploader(service)
	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(c.Store.S3Bucket),
		Key:    aws.String(fmt.Sprintf("%s.txt", c.App.Name)),
		Body:   body,
	})
	if err != nil {
		logger.LogError(err)
		return false
	}

	logger.Log(fmt.Sprintf("File uploaded to, %s", aws.StringValue(&result.Location)))
	return true
}
