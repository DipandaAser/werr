package storage

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"log"
)

type Storage struct {
	Client *s3.Client
	config StorageConfig
}

type StorageConfig struct {
	AccessKey string
	SecretKey string
	Endpoint  string
	Region    string
	Bucket    string
}

var storageClient *Storage

func InitStorage(conf StorageConfig) {
	awsCfg, err := config.LoadDefaultConfig(context.TODO(), config.WithCredentialsProvider(
		credentials.StaticCredentialsProvider{
			Value: aws.Credentials{
				AccessKeyID:     conf.AccessKey,
				SecretAccessKey: conf.SecretKey,
			},
		},
	), config.WithEndpointResolverWithOptions(aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		return aws.Endpoint{URL: conf.Endpoint}, nil
	})))

	if err != nil {
		log.Fatalln(err)
	}
	awsCfg.Region = conf.Region

	storageClient = &Storage{
		Client: s3.NewFromConfig(awsCfg),
		config: conf,
	}
}

func GetStorage() *Storage {
	return storageClient
}
