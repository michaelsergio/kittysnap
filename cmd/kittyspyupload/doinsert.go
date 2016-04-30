package main

import (
	"fmt"

	. "github.com/michaelsergio/kittyspy"
)

func main() {
	fmt.Println("Testing put")

	img_file := "test.jpg"

	uploader, err := NewMockFSUploader()
	if err != nil {
		panic(err)
	}
	res, err := uploader.Upload(img_file)
	if err != nil {
		panic(err)
	}
	fmt.Println("Upload result", res)
	fmt.Println("Tmp Dir:", uploader)
}
