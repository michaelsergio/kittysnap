package kittysnap

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
)

const (
	UploadSucess = 0
	UploadFailure
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

func (u *MockUploader) Upload(key, file string) (UploaderResult, error) {
	if u.mockError != 0 {
		return "", errors.New("error")
	}
	return "Success", nil
}

// Takes only result, no input parameters matter
type MockFSUploader struct {
	dir string
}

func NewMockFSUploader() (*MockFSUploader, error) {
	path, err := ioutil.TempDir("", "kittysnapmock")
	if err != nil {
		return nil, errors.New("NewMockFSUploader failed to create dir")
	}
	return &MockFSUploader{dir: path}, nil
}

// An absolute path input will likely break this.
func (u *MockFSUploader) Upload(file string) (UploaderResult, error) {
	// Check if file exists. If it does not, noop.
	if _, err := os.Stat(file); err != nil {
		if os.IsNotExist(err) {
			return "", errors.New("Bad Filename. Does Not Exist")
		}
	}
	// dont really copy
	newpath := filepath.Join(u.dir, file)
	abspath, abserr := filepath.Abs(file)
	if abserr != nil {
		return "", abserr
	}
	if err := os.Symlink(abspath, newpath); err != nil {
		return "", err
	}
	return "Success", nil
}

func (u *MockFSUploader) String() string {
	return u.dir
}

type MockDB struct {
	storage map[string]string
}

// In-Memory Database for Testing.
func NewMemoryDB() MockDB {
	m := make(map[string]string, 100)
	return MockDB{storage: m}
}

func (db *MockDB) PutItem(key, value string) (DatabaseResult, error) {
	db.storage[key] = value
	return DatabaseResult(key), nil
}

func (db *MockDB) GetItem(key string) (DatabaseResult, error) {
	value, ok := db.storage[key]
	if ok {
		return DatabaseResult(value), nil
	}
	return DatabaseResult(""), errors.New("Item not found")
}
