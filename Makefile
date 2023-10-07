project_name = pmgo
image_name = pmgo:pmgo

run-local:
	go run app.go

requirements:
	go mod tidy

clean-packages:
	go clean -modcache

up:
	make up-silent
	make shell

build2:
	docker build -t $(image_name) .

build-no-cache:
	docker build --no-cache -t $(image_name) .

up-silent:
	make delete-container-if-exist
	docker run -d -p 9900:9900 --name $(project_name) $(image_name) ./gapi

up-silent-prefork:
	make delete-container-if-exist
	docker run -d -p 9900:9900 --name $(project_name) $(image_name) ./gapi -prod

delete-container-if-exist:
	docker stop $(project_name) || true && docker rm $(project_name) || true

shell:
	docker exec -it $(project_name) /bin/sh

stop:
	docker stop $(project_name)

start:
	docker start $(project_name)

run-local:
	go run cmd/gapi

consul:
	docker exec -it $(project_name) /go/src/pmgo/build/consul.sh

test:
	go test -v ./...
