package main

import (
	//"path/filepath"
	"fmt"
	//"time"

	. "github.com/michaelsergio/kittysnap"
)

func main() {
	camera := NewCVCamera()
	imagefile := camera.TakeImage()
	fmt.Println(imagefile)
}
