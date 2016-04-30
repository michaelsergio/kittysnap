package kittyspy

import (
	"errors"
)

type MockCamera struct {
	imagepath string
}

func NewMockCamera(testImg string) MockCamera {
	return MockCamera{imagepath: testImg}
}

func (c *MockCamera) TakeImage() string {
	return c.imagepath
}

// Takes only result, no input parameters matter
type MockUploader struct {
	mockError int
}

func NewMockUploader(mockError int) MockUploader {
	return MockUploader{mockError: mockError}
}

func (u *MockUploader) Upload(file string) (UploaderResult, error) {
	if u.mockError != 0 {
		return "", errors.New("error")
	}
	return "Success", nil
}
