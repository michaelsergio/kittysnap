package kittysnap

import (
	"github.com/aws/aws-sdk-go/aws/credentials"
)

type Conf struct {
	camExt          string
	camLimit        int
	camOverwrite    bool
	CamDirCreated   string
	camBasename     string
	dbTable         string
	uploadBucket    string
	UploadedDir     string
	awsRegion       string
	awsCreds        *credentials.Credentials
	awsChainVerbose bool
}

func UseDefaults() Conf {
	return Conf{
		camExt:          "jpg",
		camLimit:        10,
		camOverwrite:    true,
		CamDirCreated:   "/Users/mike_sergio/.kittysnap/created",
		camBasename:     "img-",
		dbTable:         "catpics",
		uploadBucket:    "kittysnap.msergio.com",
		awsRegion:       "us-east-1",
		awsCreds:        credentials.NewEnvCredentials(),
		awsChainVerbose: true,
	}
}

// Unimplemented
func ReadFromEnv() Conf {
	conf := UseDefaults()
	/*
		if val := os.Getenv("KS_CAM_EXT"); val != "" {
			// TEST!!!
			conf.camExt = val
		}
	*/
	// TODO: Fill in
	return conf
}

// Unimplemented
func ReadFromFile() Conf {
	conf := UseDefaults()
	// TODO: Fill in.
	return conf
}
