package main

import (
	"context"
	"fmt"
	"log"
	"net/url"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
)

type myCredential struct {
	AccessKeyID     string
	SecretAccessKey string
}

func NewCredential() aws.CredentialsProvider {
	return &myCredential{
		AccessKeyID:     "trl60f",
		SecretAccessKey: "mbf9a4g",
	}
}

type myEndPoint struct {
	Region   string
	EndPoint string
}

func NewMyEndPoint() dynamodb.EndpointResolverV2 {
	return &myEndPoint{
		Region:   "us-west-2",
		EndPoint: "http://localhost:8000",
	}
}

func (m *myEndPoint) ResolveEndpoint(ctx context.Context, params dynamodb.EndpointParameters) (
	smithyendpoints.Endpoint, error,
) {
	_uri, _ := url.Parse(m.EndPoint)
	s := smithyendpoints.Endpoint{
		URI: *_uri,
	}
	return s, nil
}

func (m *myCredential) Retrieve(ctx context.Context) (aws.Credentials, error) {
	return aws.Credentials{
		AccessKeyID:     m.AccessKeyID,
		SecretAccessKey: m.SecretAccessKey,
	}, nil
}

func CreateDynamoDbConnection() {
	// Using the SDK's default configuration, loading additional config
	// and credentials values from the environment variables, shared
	// credentials, and shared configuration files
	us_west := "us-west-2"
	cfg, _ := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(us_west),
	)

	credential := NewCredential()
	endpoint := NewMyEndPoint()
	svc := dynamodb.NewFromConfig(cfg, func(o *dynamodb.Options) {
		o.Credentials = credential
		// o.BaseEndpoint = &end_point
		o.EndpointResolverV2 = endpoint
	})

	// Build the request with its input parameters
	resp, err := svc.ListTables(context.TODO(), &dynamodb.ListTablesInput{
		Limit: aws.Int32(5),
	})
	if err != nil {
		log.Fatalf("failed to list tables, %v", err)
	}

	fmt.Println("Tables:")
	for _, tableName := range resp.TableNames {
		fmt.Println(tableName)
	}
}

func main() {
	CreateDynamoDbConnection()
}
