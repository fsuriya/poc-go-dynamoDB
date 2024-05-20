package database

import (
	"errors"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

var DbClient *dynamodb.DynamoDB

func ConnectDynamo() (*dynamodb.DynamoDB, error) {
	sess, err := session.NewSession(&aws.Config{
		Endpoint: aws.String("http://localhost:8000"),
		Region:   aws.String("ap-southeast-1"),
	})
	if err != nil {
		return nil, errors.New("failed to create AWS session: " + err.Error())
	}

	log.Println("Connected to DynamoDB")
	DbClient = dynamodb.New(sess)
	return DbClient, nil
}
