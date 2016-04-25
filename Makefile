# Include AWS keys
include Keys

KEY_ENV=AWS_ACCESS_KEY_ID=$(ID) AWS_SECRET_ACCESS_KEY=$(KEY) 

BIN_NAME=kittysnap

run-mock:
	$(KEY_ENV) go run cmd/kittysnapupload/mock_main.go

run:
	$(KEY_ENV) go run cmd/kittysnapupload/main.go

run-dyn:
	$(KEY_ENV) go run cmd/kittysnapupload/doinsert.go

run-json:
	go run cmd/kittysnapupload/jsonserver.go

run-realinsert:
	$(KEY_ENV) go run cmd/kittysnapupload/realinsert.go

run-uploadsamplefile:
	$(KEY_ENV) go run cmd/kittysnapupload/upload_real.go

run-cam:
	go run cmd/kittysnapupload/realcam.go

test:
	go test -v 

fake-s3:
	fakes3 -r $(CURDIR)/mock -p 4567 

fake-dynamo:
	dynamodb-local

build:
	go build -o $(BIN_NAME) cmd/kittysnapupload/main.go
