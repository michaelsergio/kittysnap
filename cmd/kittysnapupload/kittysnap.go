package main

import (
	. "github.com/michaelsergio/kittysnap"
)

func main() {
	conf := ReadFromEnv()
	camera := NewCVCamera(&conf)
	camera.TakeImage()
}
