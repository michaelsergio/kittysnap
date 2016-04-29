package main

import (
	//"path/filepath"
	"log"
	"path/filepath"
	//"time"

	. "github.com/michaelsergio/kittysnap"
)

type Context struct {
	uploader Uploader
	camera   Camera
	db       Database
}

func main() {
	conf := ReadFromEnv()
	camera := NewCVCamera(&conf)
	uploader := NewS3Uploader(&conf)
	db := NewDynamoDatabase(&conf)
	//uploader := NewMockUploader(UploadSuccess)

	ctx := Context{uploader: &uploader, camera: &camera, db: &db}

	// TODO: Get fullpath from image
	imagefile := ctx.camera.TakeImage()

	// Get basename from file
	key := filepath.Base(imagefile)
	result, err := ctx.uploader.Upload(key, imagefile)
	if err != nil {
		log.Println("failed to upload:", err.Error())
		return
	}
	log.Println("Uploaded:", result)
	_, err = db.PutItem(imagefile, imagefile)
	if err != nil {
		log.Println("failed to put key:", err.Error())
		return
	}
	log.Println("Uploaded Key::", imagefile)
}
