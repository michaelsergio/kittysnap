package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	. "github.com/michaelsergio/kittysnap"
)

type FileUploader struct {
	uploader Uploader
	db       Database
	conf     *Conf
}

func main() {
	conf := ReadFromEnv()
	//uploader, _ := NewMockFSUploader()
	//db := NewMemoryDB()
	uploader := NewS3Uploader(&conf)
	db := NewDynamoDatabase(&conf)

	fu := FileUploader{uploader: uploader, conf: &conf, db: &db}

	filepath.Walk(conf.CamDirCreated, fu.walk)
}

func (fu *FileUploader) walk(path string, info os.FileInfo, err error) error {
	log.Println("Walking:", path)
	pathinfo, err := os.Stat(path)
	if err != nil {
		log.Println("Could not stat file:", path)
		return err
	}

	// Do nothing on directories
	if pathinfo.IsDir() {
		log.Println("Ignoring dir:", path)
		return nil
	}
	// Do nothing on dotfiles / Damn .DS_STOREs
	if strings.HasPrefix(pathinfo.Name(), ".") {
		log.Println("Ignoring dotfile:", path)
		return nil
	}

	// Ignore directories

	key := filepath.Base(path)
	result, err := fu.uploader.Upload(key, path)
	if err != nil {
		log.Println("failed to upload:", err.Error())
		return err
	}
	log.Println("Uploaded:", result)
	_, err = fu.db.PutItem(path, path)
	if err != nil {
		log.Println("failed to put key:", err.Error())
		return err
	}
	log.Println("Uploaded Key:", key)

	log.Println("Moving to upload dir: ", fu.conf.UploadedDir)
	if err := os.MkdirAll(fu.conf.UploadedDir, 0755); err != nil {
		log.Println(err)
		return err
	}
	moveTo := filepath.Join(fu.conf.UploadedDir, key)
	if err := os.Rename(path, moveTo); err != nil {
		log.Println(err)
		return err
	}
	log.Println("Moved to uploaded dir")

	return nil
}
