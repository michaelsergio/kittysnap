package main

import (
	//"path/filepath"
	"log"
	//"time"

	. "github.com/michaelsergio/kittysnap"
)

func main() {
	conf := UseDefaults()
	uploader := NewS3Uploader(&conf)

	imagefile := "test.jpg"
	result, err := uploader.Upload("sample1", imagefile)
	if err != nil {
		log.Println("failed to upload:", err.Error())
		return
	}
	log.Println(result)
}
