.PHONY: build
build:
	GO111MODULE=on go build ./cmd/server

.PHONY: docker-start
docker-start:
	docker-compose up --build --detach

.PHONY: server-start
server-start:
	./server
