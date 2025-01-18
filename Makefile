ROOT_DIR=/Users/k
# ROOT_DIR=C:\

run:
	cd cmd && go run main.go -root $(ROOT_DIR)

build:
	go mod vendor
	docker build --platform linux/amd64 -t anosaa/file-map-amd64:0.0.2 .
	# docker build --platform linux/arm64 -t anosaa/file-map-arm64:0.0.2 .
	rm -rf vendor