package kittysnap

import (
	"encoding/json"
)

type Lister struct{}

func ListImagesJson(lister *Lister) ([]byte, error) {
	type Example struct {
		S string
		I int
	}
	items := Example{S: "hello", I: 5}
	return json.Marshal(items)
}
