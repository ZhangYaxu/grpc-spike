build_go:
	mkdir -p recording
	protoc -I /usr/include --proto_path=proto --go_out=plugins=grpc:./recording proto/*.proto
	go build -o server ./cmd/server
	go build -o client ./cmd/client

build_php:
	mkdir -p php
	protoc -I /usr/include --proto_path=proto --php_out=php/  --grpc_out=class_suffix=:php/ --plugin=protoc-gen-grpc=/home/rodrigodiez/grpc/bins/opt/grpc_php_plugin proto/*.proto
