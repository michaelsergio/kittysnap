package main

import (
	//"log"
	. "github.com/michaelsergio/kittysnap"
)

func main() {
	rq, err := OpenRedisUploadQueue()

	rq.AppendCreatedFile("hello/world")

	rq.CloseRedisUploadQueue()

}
