package kittyspy

type Camera interface {
	TakeImage() string
}

type UploaderResult string

// Add a string type for UploaderResult

type Uploader interface {
	Upload(file string) (UploaderResult, error)
}
