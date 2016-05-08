package main

import (
	. "github.com/michaelsergio/kittysnap"
	//"log"
)

func main() {
	conf := ReadFromEnv()
	camera := NewCVCamera(&conf)
	camera.TakeImage()
	/*
		path, err := camera.TakeImage()
		if err != nil {
			log.Println("Error taking image:", err)
		} else {
			log.Println("Image saved to", path)
		}
	*/

}
