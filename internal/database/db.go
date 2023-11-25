package database

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func NewDBClient() *dynamodb.Client {
	awsRegion := os.Getenv("AWS_REGION")

	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(awsRegion))
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}
	return dynamodb.NewFromConfig(cfg)
}
