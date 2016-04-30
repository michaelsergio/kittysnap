# Include AWS keys
include Keys

KEY_ENV=AWS_ACCESS_KEY_ID=$(ID) AWS_SECRET_ACCESS_KEY=$(KEY) 

BIN_NAME=kittysnap

run-mock:
	$(KEY_ENV) go run cmd/kittyspyupload/mock_main.go

run:
	$(KEY_ENV) go run cmd/kittyspyupload/main.go

run-dyn:
	$(KEY_ENV) go run cmd/kittyspyupload/doinsert.go

run-json:
	go run cmd/kittyspyupload/jsonserver.go

run-realinsert:
	$(KEY_ENV) go run cmd/kittyspyupload/realinsert.go

run-uploadsamplefile:
	$(KEY_ENV) go run cmd/kittyspyupload/upload_real.go

run-cam:
	go run cmd/kittyspyupload/realcam.go

test:
	go test -v 

fake-s3:
	fakes3 -r $(CURDIR)/mock -p 4567 

fake-dynamo:
	dynamodb-local

build:
	go build -o $(BIN_NAME) cmd/kittyspyupload/main.go
