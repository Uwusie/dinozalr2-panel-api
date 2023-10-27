package database

import (
	"context"
	"log"
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
		cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("eu-north-1"))
		if err != nil {
			log.Fatalf("unable to load SDK config, %v", err)
		}
		instance = dynamodb.NewFromConfig(cfg)
	})
	return instance
}
