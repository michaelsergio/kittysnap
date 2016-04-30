package kittysnap

import (
	"github.com/aws/aws-sdk-go/aws/credentials"
)

type Conf struct {
	camExt          string
	camLimit        int
	camOverwrite    bool
	camDir          string
	camBasename     string
	dbTable         string
	uploadBucket    string
	awsRegion       string
	awsCreds        *credentials.Credentials
	awsChainVerbose bool
}

func UseDefaults() Conf {
	return Conf{
		camExt:          "jpg",
		camLimit:        10,
		camOverwrite:    true,
		camDir:          "~/.kittyspy",
		camBasename:     "img-",
		dbTable:         "catpics",
		uploadBucket:    "kittysnap",
		awsRegion:       "us-east-1",
		awsCreds:        credentials.NewEnvCredentials(),
		awsChainVerbose: true,
	}
}

// Unimplemented
func ReadFromEnv() Conf {
	conf := UseDefaults()
	// TODO: Fill in
	return conf
}

// Unimplemented
func ReadFromFile() Conf {
	conf := UseDefaults()
	// TODO: Fill in.
	return conf
}
