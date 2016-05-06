package kittysnap

// Add entries to keystore.

import (
	"fmt"
	"log"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

const DbZTime = "2006-01-02T15:04:05Z"

type DynamoDatabase struct {
	tablename string
	location  string
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
	// The value should be the s3 location
	s3path := value

	// The key MUST be a filename in the form  123124.jpg

	// Pull time out of filename as unix time
	unixtimeStr := strings.TrimSuffix(key, filepath.Ext(key))
	unixSec, err := strconv.Atoi(unixtimeStr)

	taken := time.Unix(int64(unixSec), 0)
	yyyy, ww := taken.ISOWeek()

	yearweek := fmt.Sprintf("%s-%s", yyyy, ww)
	//curtime := time.Now().Format(DbZTime)
	//log.Println(curtimeI)
	// TODO: Fix parameters
	params := &dynamodb.PutItemInput{
		Item: map[string]*dynamodb.AttributeValue{
			"year-weeknum": {
				S: aws.String(yearweek),
			},
			"unixtime": {
				N: aws.String(unixtimeStr),
			},
			"s3url": {
				S: aws.String(s3path),
			},
			"location": {
				S: aws.String(db.location),
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
