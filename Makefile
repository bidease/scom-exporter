linux:
	GOOS=linux GOARCH=amd64 go build -o scom-exporter-linux-amd64 ./

darwin:
	GOOS=darwin GOARCH=amd64 go build -o scom-exporter-darwin-amd64 ./

.PHONY: linux darwin
