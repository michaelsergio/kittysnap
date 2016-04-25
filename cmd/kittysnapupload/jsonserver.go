package main

import (
	//"fmt"
	"os"

	. "github.com/michaelsergio/kittysnap"
)

func main() {
	serialized, err := ListImagesJson(nil)
	if err != nil {
		panic(err)
	}
	//fmt.Println(serialized)
	os.Stdout.Write(serialized)
}
