package auth_service

import (
	"fmt"
	"poc-go-dynamodb/database"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func ExistsUser(username string) (bool, error) {
	fmt.Println("Checking if user exists")
	fmt.Println("Username: ", username)
	result, err := database.DbClient.GetItem(&dynamodb.GetItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"username": {
				S: aws.String(username),
			},
		},
		TableName: aws.String("user"),
	})

	if err != nil {
		return false, err
	}

	if result.Item == nil {
		return false, nil
	}
	return true, nil
}
