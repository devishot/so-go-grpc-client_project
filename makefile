
run:
	docker-compose -f build/dev/docker-compose.yml up

test:
	go generate ./...
	go test ./...

generate_mock:
	go generate ./...

generate_proto:
	protoc -I=interfaces/grpc/protofiles -I=${GOPATH}/src --go_out=plugins=grpc:interfaces/grpc/api/ client.proto client_project.proto connection.proto
