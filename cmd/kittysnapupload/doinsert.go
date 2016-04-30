package main

import (
	"log"

	. "github.com/michaelsergio/kittysnap"
)

func main() {
	log.Println("Testing put")

	img_file := "test.jpg"

	uploader, err := NewMockFSUploader()
	if err != nil {
		panic(err)
	}
	res, err := uploader.Upload(img_file)
	if err != nil {
		panic(err)
	}
	log.Println("Upload result", res)
	log.Println("Tmp Dir:", uploader)
}
