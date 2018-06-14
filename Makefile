GOPATH=/go
VOLUME_ROOT=/app
CODEGEN_ROOT=${VOLUME_ROOT}
DOCS_ROOT=${CODEGEN_ROOT}/docs
GRPC_WD=${VOLUME_ROOT}
PROTOC=docker-compose run --rm protoc
PROTOC_FILE=${VOLUME_ROOT}/gwsamplepb/grpc-gw-sample.proto

build: codegens dep docker

codegens: code doc gateway swagger

docker:
	docker-compose build

test:
	docker-compose run --rm grpc go test -v -cover ./...

dep:
	docker-compose run --rm dep ensure -v

code:
	${PROTOC} -I${GRPC_WD} \
	-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
	--go_out=plugins=grpc:${CODEGEN_ROOT} \
	${PROTOC_FILE}

doc:
	${PROTOC} -I${GRPC_WD} \
	-I${GOPATH}/src/github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc \
	-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
	--doc_out=${DOCS_ROOT} --doc_opt=html,application.html \
	${PROTOC_FILE}

gateway:
	${PROTOC} -I${GRPC_WD} \
	-I$(GOPATH)/src \
	-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
	--grpc-gateway_out=logtostderr=true:${CODEGEN_ROOT} \
	${PROTOC_FILE}

swagger:
	${PROTOC} -I${GRPC_WD} \
	-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
	--swagger_out=logtostderr=true:${CODEGEN_ROOT} \
	${PROTOC_FILE}
