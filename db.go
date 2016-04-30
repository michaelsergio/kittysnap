package kittyspy

// Add entries to keystore.

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type DynamoDatabase struct {
	tablename string
	dyndb     *dynamodb.DynamoDB
}

func NewDynamoDatabase(awsconf *aws.Config) DynamoDatabase {
	return DynamoDatabase{
		tablename: "catpics",
		dyndb:     dynamodb.New(session.New(), awsconf),
	}

}

func (db *DynamoDatabase) PutItem(key, value string) (DatabaseResult, error) {
	params := &dynamodb.PutItemInput{
		Item: map[string]*dynamodb.AttributeValue{
			/*
				key: {
					S: aws.String(value),
				},
			*/
			"filename": {
				S: aws.String(key),
			},
			"s3id": {
				S: aws.String(value),
			},
			"date": {
				S: aws.String(value),
			},
		},
		TableName: aws.String(db.tablename), // Required
	}

	resp, err := db.dyndb.PutItem(params)
	if err != nil {
		return "", err
	}
	fmt.Println("Dynamo Response:", resp)
	return "", err
}

func (db *DynamoDatabase) GetItem(key string) (DatabaseResult, error) {
	return "", nil
}
