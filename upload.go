package kittyspy

import (
	"fmt"
	"log"
	"strings" // replace with io reader later.

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func ListContents() {

}

func Upload() {
	fmt.Println("Uploading....")

	//svc := s3.New(session.New(), &aws.Config{Region: aws.String("mock-region")})
	creds := credentials.NewEnvCredentials()
	config := aws.NewConfig().
		WithRegion("us-east-1").
		WithCredentials(creds).
		WithCredentialsChainVerboseErrors(true)
	// config := &aws.Config{Region: aws.String("us-west-2")}
	// config = config.WithCredentialsChainVerboseErrors(true)
	svc := s3.New(session.New(), config)

	// List buckets
	result, err := svc.ListBuckets(&s3.ListBucketsInput{})
	if err != nil {
		log.Println("Failed to list buckets", err)
		return
	}

	log.Println("Buckets:")
	for _, bucket := range result.Buckets {
		log.Printf("%s : %s\n", aws.StringValue(bucket.Name), bucket.CreationDate)
	}

	// Create Bucket
	// Do not want
	/*
		result, err := svc.CreateBucket(&s3.CreateBucketInput{
			Bucket: &bucket,
		})
		if err != nil {
			log.Println("Failed to create bucket", err)
			return
		}

		if err = svc.WaitUntilBucketExists(&s3.HeadBucketInput{Bucket: &bucket}); err != nil {
			log.Printf("Failed to wait for bucket to exist %s, %s\n", bucket, err)
			return
		}

	*/

	// TODO: define Bucket name
	// Replace Body with file IO reader.
	// TODO define keyname
	bucket := "kittyspy"
	key := "test"
	uploadResult, err := svc.PutObject(&s3.PutObjectInput{
		Body:   strings.NewReader("Hello World!"),
		Bucket: &bucket,
		Key:    &key,
	})
	if err != nil {
		log.Printf("Failed to upload data to %s/%s, %s\n", bucket, key, err.Error())
		return
	}

	log.Printf("Upload Result", uploadResult)
	//log.Printf("Successfully created bucket %s and uploaded data with key %s\n", bucket, key)

	//log.Printf("Uploaded to ", uploadResult.Location(), "ver:", uploadResult.VersionID, "id:", uploadResult.UploadID)

}
