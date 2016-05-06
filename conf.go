package kittysnap

import (
	"github.com/aws/aws-sdk-go/aws/credentials"
	"os"
	"path/filepath"
)

type Conf struct {
	camExt          string
	camLimit        int
	camOverwrite    bool
	CamDirCreated   string
	camBasename     string
	dbTable         string
	location        string
	uploadBucket    string
	UploadedDir     string
	awsRegion       string
	awsCreds        *credentials.Credentials
	awsChainVerbose bool
}

func UseDefaults() Conf {
	home := os.Getenv("HOME")
	workingDir := filepath.Join(home, ".kittysnap")

	return Conf{
		CamDirCreated:   filepath.Join(workingDir, "created"),
		UploadedDir:     filepath.Join(workingDir, "uploaded"),
		awsChainVerbose: true,
		awsCreds:        credentials.NewEnvCredentials(),
		awsRegion:       "us-east-1",
		camBasename:     "img-",
		camExt:          "jpg",
		camLimit:        10,
		camOverwrite:    true,
		dbTable:         "catpics",
		location:        "UnknownLocation",
		uploadBucket:    "kittysnap.msergio.com",
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
