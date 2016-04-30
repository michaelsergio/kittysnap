package main

import (
	//"path/filepath"
	"log"
	//"time"

	. "github.com/michaelsergio/kittysnap"
)

func main() {
	conf := UseDefaults()
	camera := NewCVCamera(&conf)
	imagefile := camera.TakeImage()
	log.Println(imagefile)
}
