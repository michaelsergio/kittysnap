package main

import (
	//"path/filepath"
	"fmt"
	//"time"

	. "github.com/michaelsergio/kittysnap"
)

func main() {
	conf := UseDefaults()
	camera := NewCVCamera(&conf)
	imagefile := camera.TakeImage()
	fmt.Println(imagefile)
}
