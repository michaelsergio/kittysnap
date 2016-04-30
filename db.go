package kittyspy

// Add entries to keystore.

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func PutImageEntry(value string) {
	tablename := "kittyspy"
	key := "imagename"

	creds := credentials.NewEnvCredentials()
	config := aws.NewConfig().
		WithRegion("us-east-1").
		WithCredentials(creds).
		WithCredentialsChainVerboseErrors(true)
	dyndb := dynamodb.New(session.New(), config)

	params := &dynamodb.PutItemInput{
		Item: map[string]*dynamodb.AttributeValue{
			key: {
				S: aws.String(value),
			},
		},
		TableName: aws.String(tablename), // Required
	}

	resp, err := dyndb.PutItem(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)

}
