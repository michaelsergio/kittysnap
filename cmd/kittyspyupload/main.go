package main

import (
	//"path/filepath"
	"fmt"
	//"time"

	. "github.com/michaelsergio/kittyspy"
)

type Context struct {
	uploader Uploader
	camera   Camera
}

func RunEverySeconds(seconds uint32, f func()) {
	// Write image to tmp storage
	//timer := time.NewTimer(time.Duration(seconds) * time.Second)
	// TODO: Timer currently does not tick more than once.
	f()

}

func main() {
	camera := NewMockCamera("test.jpg")
	uploader := NewMockUploader(0)
	//uploader := NewMockUploader(UploadSuccess)

	ctx := Context{uploader: &uploader, camera: &camera}

	RunEverySeconds(10, func() {
		imagefile := ctx.camera.TakeImage()
		result, err := ctx.uploader.Upload(imagefile)
		if err != nil {
			fmt.Println("failed to upload:", err.Error())
			return
		}
		fmt.Println("Uploaded:", result)
	})
}
