package kittysnap

import (
	"fmt"
	//"io/ioutil"
	"path"
	// "time"

	"github.com/lazywei/go-opencv/opencv"
)

type CVCamera struct {
	ext         string
	limit       int
	overwrite   bool
	dir         string
	base        string
	imagesTaken int
}

func NewCVCamera() CVCamera {
	/*
		tmp, err := ioutil.TempDir("", "kittysnap")
		if err != nil {
			fmt.Println("Could not get temp directory")
			tmp = ""
		}
	*/
	tmp := "/tmp/kittysnapimages"
	return CVCamera{
		ext:         "jpg",
		limit:       10, // Take time for lighting and focus
		overwrite:   true,
		dir:         tmp,
		base:        "img-",
		imagesTaken: 0,
	}
}

func (cam *CVCamera) TakeImage() string {
	capture := opencv.NewCameraCapture(0)
	if capture == nil {
		return "error"
	}

	filename := "error"

	// Take several frames of images
	for i := 0; i < cam.limit; i++ {
		img := capture.QueryFrame()
		cam.imagesTaken += 1

		if cam.overwrite {
			// only write last image
			if i == cam.limit-1 {
				filename = cam.writeImage(img, 0)
			}
		} else {
			// Write every image
			filename = cam.writeImage(img, cam.imagesTaken)
		}
	}
	return filename
}

// Handle Errors in a non sloppy manner.

// Returns path to written image.
func (cam *CVCamera) writeImage(img *opencv.IplImage, id int) string {
	name := fmt.Sprintf("%s%05d.%s", cam.base, id, cam.ext)
	//timestamp := strconv.FormatInt(time.Now().UTC().UnixNano(), 10)
	//name := fmt.Sprintf("%s.%s", timestamp, ext)
	filename := path.Join(cam.dir, name)

	// Must test failing to save (use bad ext)
	rv := opencv.SaveImage(filename, img, 0)
	if rv != 1 {
		fmt.Println("Failed to save image:", filename, "code", rv)
		return ""
		//return "error"
	} else {
		id++
	}
	return filename
}
