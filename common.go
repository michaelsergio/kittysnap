package kittyspy

type Camera interface {
	TakeImage() string
}

type UploaderResult string

// Add a string type for UploaderResult

type Uploader interface {
	Upload(file string) (UploaderResult, error)
}

type DatabaseResult string
type Database interface {
	PutItem(key, value string) (DatabaseResult, error)
	GetItem(key string) (DatabaseResult, error)
}
