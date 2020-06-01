package config

import (
	"fmt"
	"os"
	"testing"
)

func TestGet(t *testing.T) {
	version := "42.42.42"
	clientID := "user42"
	clientSecret := "secret42"
	accessToken := "token42"
	awsRegion := "us-east-1"

	os.Setenv("VERSION", version)
	os.Setenv("CLIENT_ID", clientID)
	os.Setenv("CLIENT_SECRET", clientSecret)
	os.Setenv("ACCESS_TOKEN", accessToken)
	os.Setenv("AWS_REGION", awsRegion)

	config := Get()

	errorMessage := fmt.Sprintf(
		"CASE 1: Should be able to process environment variables (VERSION = %s, CLIENT_ID = %s, CLIENT_SECRET = %s, ACCESS_TOKEN = %s, AWS_REGION = %s)",
		version, clientID, clientSecret, accessToken, awsRegion,
	)

	if config.App.Name != "memoriesbot" ||
		config.App.Version != version ||
		config.Auth.GooglePhotos.ClientID != clientID ||
		config.Auth.GooglePhotos.ClientSecret != clientSecret ||
		config.Auth.GooglePhotos.AccessToken != accessToken ||
		config.Aws.Region != awsRegion {
		t.Error(errorMessage)
	}
}
