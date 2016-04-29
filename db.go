package kittysnap

// Add entries to keystore.

import (
	"log"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

const DbZTime = "2006-01-02T15:04:05Z"

type DynamoDatabase struct {
	tablename string
	dyndb     *dynamodb.DynamoDB
}

func NewDynamoDatabase(conf *Conf) DynamoDatabase {
	awsconf := aws.NewConfig().
		WithRegion(conf.awsRegion).
		WithCredentials(conf.awsCreds).
		WithCredentialsChainVerboseErrors(conf.awsChainVerbose)

	return DynamoDatabase{
		tablename: conf.dbTable,
		dyndb:     dynamodb.New(session.New(), awsconf),
	}

}

func (db *DynamoDatabase) PutItem(key, value string) (DatabaseResult, error) {

	//curtime := time.Now().Format(time.RFC3339)
	curtime := time.Now().Format(DbZTime)
	log.Println(curtime)
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
				S: aws.String(curtime),
			},
		},
		TableName: aws.String(db.tablename), // Required
	}

	resp, err := db.dyndb.PutItem(params)
	if err != nil {
		return "", err
	}
	log.Println("Dynamo Response:", resp)
	return "", err
}

func (db *DynamoDatabase) GetItem(key string) (DatabaseResult, error) {
	return "", nil
}
