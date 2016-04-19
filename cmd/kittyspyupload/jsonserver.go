package main

import (
	//"fmt"
	"os"

	. "github.com/michaelsergio/kittyspy"
)

func main() {
	serialized, err := ListImagesJson(nil)
	if err != nil {
		panic(err)
	}
	//fmt.Println(serialized)
	os.Stdout.Write(serialized)
}
