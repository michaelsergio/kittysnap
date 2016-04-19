package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"

	. "github.com/michaelsergio/kittyspy"
)

func main() {
	fmt.Println("Testing Dynamo Insert")

	creds := credentials.NewEnvCredentials()
	config := aws.NewConfig().
		WithRegion("us-east-1").
		WithCredentials(creds).
		WithCredentialsChainVerboseErrors(true)

	key := "imagename"
	db := NewDynamoDatabase(config)
	resp, err := db.PutItem(key, "world")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(resp)
}
