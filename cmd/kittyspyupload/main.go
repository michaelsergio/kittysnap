package main

import (
	//"path/filepath"
	"fmt"
	"path/filepath"
	//"time"

	. "github.com/michaelsergio/kittyspy"
)

type Context struct {
	uploader Uploader
	camera   Camera
	db       Database
}

func main() {
	camera := NewMockCamera("test.jpg")
	uploader := NewMockUploader(0)
	db := NewMemoryDB()
	//uploader := NewMockUploader(UploadSuccess)

	ctx := Context{uploader: &uploader, camera: &camera, db: &db}

	// TODO: Get fullpath from image
	imagefile := ctx.camera.TakeImage()

	// Get basename from file
	key := filepath.Base(imagefile)
	result, err := ctx.uploader.Upload(key, imagefile)
	if err != nil {
		fmt.Println("failed to upload:", err.Error())
		return
	}
	fmt.Println("Uploaded:", result)
	_, err = db.PutItem(imagefile, imagefile)
	if err != nil {
		fmt.Println("failed to put key:", err.Error())
		return
	}
	fmt.Println("Uploaded Key::", imagefile)
}
