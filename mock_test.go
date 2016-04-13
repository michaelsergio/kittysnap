package kittyspy

import (
	"testing"
)

func TestMockUploadSuccess(t *testing.T) {
	uploader := NewMockUploader(0)
	_, err := uploader.Upload("test.jpg")
	if err != nil {
		t.Fail()
	}
}
