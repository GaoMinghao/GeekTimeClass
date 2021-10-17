export tag=v1.0
root:
	export ROOT=GeekTimeClass.cloud_native_campus.homework

build:
	echo "building httpserver binary"
	mkdir -p bin/amd64
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/amd64 .

release: build
	echo "building httpserver container"
	docker build -t michaelindeui/httpserver:${tag} .

push: release
	echo "pushing michaelindeui/httpserver"
	docker push michaelindeui/httpserver:v1.0
