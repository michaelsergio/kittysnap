# Include AWS keys
include Keys

KEY_ENV=AWS_ACCESS_KEY_ID=$(ID) AWS_SECRET_ACCESS_KEY=$(KEY) 

SNAP_BIN_NAME=kittysnap
UPLOAD_BIN_NAME=kittysnapuploader

run-mock:
	$(KEY_ENV) go run cmd/kittysnapupload/mock_main.go

snap:
	$(KEY_ENV) go run cmd/kittysnapupload/kittysnap.go

uploader:
	$(KEY_ENV) go run cmd/kittysnapupload/kittysnapuploader.go

run-allinone:
	$(KEY_ENV) go run cmd/kittysnapupload/allinone.go

try-redis:
	go run cmd/kittysnapupload/tryredis.go

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
	go build -o $(SNAP_BIN_NAME) cmd/kittysnapupload/kittysnap.go
	#go build -o $(UPLOAD_BIN_NAME) cmd/kittysnapupload/main.go
