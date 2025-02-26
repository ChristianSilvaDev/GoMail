package container

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/aws/session"
)

func ProvideDynamoDB() *dynamodb.DynamoDB {
	return dynamodb.New(session.Must(session.NewSession()))
}
