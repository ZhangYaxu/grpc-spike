dependencies:
	dep ensure

proto: dependencies
	rm -fr recording
	mkdir -p recording
	protoc \
		-I/usr/include \
		-I./vendor \
		-I./vendor/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--proto_path=proto \
		--go_out=plugins=grpc:./recording proto/recording.proto \
		--grpc-gateway_out=logtostderr=true:./recording \
		--swagger_out=logtostderr=true:./recording

go: dependencies proto
	rm -fr bin
	mkdir -p bin/linux
	go build -o bin/client ./cmd/client
	GOOS=linux GOARCH=amd64 go build -o bin/linux/server ./cmd/server
	GOOS=linux GOARCH=amd64 go build -o bin/linux/gw ./cmd/gw

php: dependencies
	rm -fr php/src
	mkdir -p php/src
	protoc \
	-I/usr/include \
	-I./vendor \
	-I./vendor/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
	--proto_path=proto \
	--php_out=php/src  \
	--grpc_out=php/src \
	--plugin=protoc-gen-grpc=/Users/rodrigo.diez/grpc/bins/opt/grpc_php_plugin proto/recording.proto

	protoc \
	-I/usr/include \
	-I./vendor \
	-I./vendor/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
	--proto_path=proto \
	--php_out=php/src  \
	--grpc_out=php/src \
	--plugin=protoc-gen-grpc=/Users/rodrigo.diez/grpc/bins/opt/grpc_php_plugin vendor/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api/annotations.proto
