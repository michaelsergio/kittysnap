package kittysnap

import (
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func ListContents() {

}

type S3Uploader struct {
	svc    *s3.S3
	bucket string
}

func NewS3Uploader(conf *Conf) S3Uploader {
	config := aws.NewConfig().
		WithRegion(conf.awsRegion).
		WithCredentials(conf.awsCreds).
		WithCredentialsChainVerboseErrors(conf.awsChainVerbose)
	uploader := S3Uploader{
		svc:    s3.New(session.New(), config),
		bucket: conf.uploadBucket,
	}
	return uploader
}

func (up *S3Uploader) Upload(key, filename string) (UploaderResult, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}

	log.Println("Uploading", filename, "to s3 with key", key)

	uploadResult, err := up.svc.PutObject(&s3.PutObjectInput{
		Body:   file,
		Bucket: &up.bucket,
		Key:    &key,
	})
	if err != nil {
		return "", err
	}

	upRes := UploaderResult(uploadResult.String())
	return upRes, nil
}
