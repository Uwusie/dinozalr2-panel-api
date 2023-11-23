package database

import (
	"context"
	"log"
	"os"
	"sync"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

var (
	instance *dynamodb.Client
	once     sync.Once
)

func GetDBClient() *dynamodb.Client {
	once.Do(func() {
		awsRegion := os.Getenv("AWS_REGION")

		cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(awsRegion))
		if err != nil {
			log.Fatalf("unable to load SDK config, %v", err)
		}
		instance = dynamodb.NewFromConfig(cfg)
	})
	return instance
}
