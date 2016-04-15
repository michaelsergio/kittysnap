package kittyspy

import (
	"testing"
)

func TestMockUploadSuccess(t *testing.T) {
	uploader := NewMockUploader(0)
	//uploader := NewMockUploader(UploadSuccess)
	_, err := uploader.Upload("test.jpg")
	if err != nil {
		t.Fail()
	}
}

func TestMockError(t *testing.T) {
	uploader := NewMockUploader(1)
	//uploader := NewMockUploader(UploadFail)
	_, err := uploader.Upload("test.jpg")
	if err == nil {
		t.Fail()
	}
}
