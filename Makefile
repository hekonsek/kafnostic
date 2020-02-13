build:
	GO111MODULE=on go build -o kafnostic main/root.go main/produce.go main/consume.go

docker-build: build
	docker build . -t hekonsek/kafnostic

docker-push: docker-build
	docker push hekonsek/kafnostic