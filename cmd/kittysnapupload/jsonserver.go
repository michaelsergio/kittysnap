package main

import (
	//"log"
	"os"

	. "github.com/michaelsergio/kittysnap"
)

func main() {
	serialized, err := ListImagesJson(nil)
	if err != nil {
		panic(err)
	}
	//log.Println(serialized)
	os.Stdout.Write(serialized)
}
