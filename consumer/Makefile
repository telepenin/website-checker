.EXPORT_ALL_VARIABLES:

CONFIG=../example/config.yaml

# Run go fmt against code
fmt:
	go fmt ./...

# Run go vet against code
vet:
	go vet ./...

run: fmt vet
	go run main.go