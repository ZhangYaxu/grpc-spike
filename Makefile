GRPC_PATH = /home/rodrigodiez/src/github.com/grpc/grpc
PROTOC = $(GRPC_PATH)/bins/opt/protobuf/protoc
GRPC_PHP_PLUGIN = $(GRPC_PATH)/bins/opt/grpc_php_plugin

dependencies:
	dep ensure

proto: dependencies
	rm -fr recording
	mkdir -p recording
	$(PROTOC) \
		-I/usr/include \
		-I./vendor \
		-I./proto \
		-I./vendor/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--go_out=plugins=grpc:./recording \
		--grpc-gateway_out=logtostderr=true:./recording \
		--swagger_out=logtostderr=true:./recording \
		proto/recording.proto

go: dependencies proto
	rm -fr bin
	mkdir -p bin/linux
	go build -o bin/client ./cmd/client
	GOOS=linux GOARCH=amd64 go build -o bin/linux/server ./cmd/server
	GOOS=linux GOARCH=amd64 go build -o bin/linux/gw ./cmd/gw

php: dependencies
	rm -fr php/src
	mkdir -p php/src
	$(PROTOC) \
	-I/usr/include \
	-I./vendor \
	-I./proto \
	-I./vendor/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
	--php_out=php/src  \
	--grpc_out=php/src \
	--plugin=protoc-gen-grpc=$(GRPC_PHP_PLUGIN) proto/recording.proto

	protoc \
	-I/usr/include \
	-I./vendor/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
	--php_out=php/src  \
	--grpc_out=php/src \
	--plugin=protoc-gen-grpc=$(GRPC_PHP_PLUGIN) vendor/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api/*.proto
