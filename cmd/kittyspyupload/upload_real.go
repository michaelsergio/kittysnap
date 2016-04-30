package main

import (
	//"path/filepath"
	"fmt"
	//"time"

	. "github.com/michaelsergio/kittyspy"
)

func main() {
	uploader := NewS3Uploader()

	imagefile := "test.jpg"
	result, err := uploader.Upload("sample1", imagefile)
	if err != nil {
		fmt.Println("failed to upload:", err.Error())
		return
	}
	fmt.Println(result)
}
