package kittysnap

import (
	"testing"
)

func TestMockUploadSuccess(t *testing.T) {
	uploader := NewMockUploader(0)
	//uploader := NewMockUploader(UploadSuccess)
	_, err := uploader.Upload("test.jpg")
	if err != nil {
		t.Error()
	}
}

func TestMockError(t *testing.T) {
	uploader := NewMockUploader(1)
	//uploader := NewMockUploader(UploadError)
	_, err := uploader.Upload("test.jpg")
	if err == nil {
		t.Error()
	}
}

func TestMockDatabase(t *testing.T) {
	db := NewMemoryDB()
	res, err := db.PutItem("key", "value")
	if err != nil {
		t.Error("key1 put err:", err)
	}
	if string(res) != "key" {
		t.Error("key1 put val:", string(res))
	}
	res, err = db.PutItem("key2", "value2")
	if err != nil {
		t.Error("key2 put err:", err)
	}
	if string(res) != "key2" {
		t.Error("key2 put val:")
	}
	res, err = db.GetItem("key")
	if err != nil {
		t.Error("key1 get err:")
	}
	if string(res) != "value" {
		t.Error("key1 get val:")
	}
	res, err = db.GetItem("key2")
	if err != nil {
		t.Error("key2 get err:")
	}
	if string(res) != "value2" {
		t.Error("key2 get val:", string(res))
	}

	res, err = db.GetItem("key3")
	if err == nil {
		t.Error("key 3 get err:")
	}

	res, err = db.PutItem("key2", "value2update")
	if err != nil {
		t.Error("key2 put update:")
	}
	if string(res) != "key2" {
		t.Error("key2 update val:")
	}
	res, err = db.GetItem("key2")
	if err != nil {
		t.Error("key2 update err:")
	}
	if string(res) != "value2update" {
		t.Error("key2 update val:")
	}

}
