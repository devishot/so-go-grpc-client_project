
run:
	docker-compose -f build/dev/docker-compose.yml up

test:
	go generate ./...
	go test ./...