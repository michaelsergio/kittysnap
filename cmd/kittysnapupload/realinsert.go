package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"

	. "github.com/michaelsergio/kittysnap"
)

func main() {
	// Seems to be broken
	fmt.Println("Testing Dynamo Insert")

	creds := credentials.NewEnvCredentials()
	config := aws.NewConfig().
		WithRegion("us-west-2").
		WithCredentials(creds).
		WithCredentialsChainVerboseErrors(true)

	key := "imagename"
	db := NewDynamoDatabase(config)
	resp, err := db.PutItem(key, "world")
	if err != nil {
		fmt.Println("Put Error:", err)
		return
	}
	fmt.Println(resp)
}
