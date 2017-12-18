GRPC_PATH = $(GOPATH)/src/github.com/grpc/grpc
PROTOC = $(GRPC_PATH)/bins/opt/protobuf/protoc
GRPC_PHP_PLUGIN = $(GRPC_PATH)/bins/opt/grpc_php_plugin

dependencies:
	dep ensure

proto: dependencies
	mkdir -p recording
	mkdir -p nodejs/proto

	$(PROTOC) \
		-I./include \
		-I./vendor \
		-I./proto \
		-I./vendor/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--go_out=plugins=grpc:./recording \
		--grpc-gateway_out=logtostderr=true:./recording \
		--swagger_out=logtostderr=true:./recording \
		proto/recording.proto

		cat proto/recording.proto | sed '/import \"google\/api\/annotations\.proto\";/ d' > nodejs/proto/recording.proto

go: dependencies proto
	mkdir -p bin/linux
	go build -o bin/client ./cmd/client
	GOOS=linux GOARCH=amd64 go build -o bin/linux/server ./cmd/server
	GOOS=linux GOARCH=amd64 go build -o bin/linux/gw ./cmd/gw

php: dependencies
	mkdir -p php/src
	$(PROTOC) \
	-I./include \
	-I./vendor \
	-I./proto \
	-I./vendor/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
	--php_out=php/src  \
	--grpc_out=php/src \
	--plugin=protoc-gen-grpc=$(GRPC_PHP_PLUGIN) proto/recording.proto

	$(PROTOC) \
	-I./include \
	-I./vendor/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
	--php_out=php/src  \
	--grpc_out=php/src \
	--plugin=protoc-gen-grpc=$(GRPC_PHP_PLUGIN) vendor/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api/*.proto

php_plugin:
	cd $(GOPATH)/src/github.com/grpc/grpc && git checkout v1.8.1 &&	git submodule update --init &&	make grpc_php_plugin

docker:
	docker build -t kobaltmusic/grpc-spike:php -f php/Dockerfile ./php
	docker build -t kobaltmusic/grpc-spike:node -f nodejs/Dockerfile ./nodejs

clean:
	rm -fr recording
	rm -fr nodejs/proto
	rm -fr bin
	rm -fr php/src
	rm -fr vendor


all: clean php_plugin proto go php docker
