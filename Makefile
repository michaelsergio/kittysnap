# Include AWS keys
include Keys

run:
	AWS_ACCESS_KEY_ID=$(ID) AWS_SECRET_ACCESS_KEY=$(KEY) go run cmd/kittyspyupload/main.go

test:
	go test -v 

