package main

import (
	//"path/filepath"
	"fmt"
	//"time"

	. "github.com/michaelsergio/kittyspy"
)

func main() {
	camera := NewCVCamera()
	imagefile := camera.TakeImage()
	fmt.Println(imagefile)
}
