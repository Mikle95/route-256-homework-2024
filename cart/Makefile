LOCAL_BIN:=$(CURDIR)/../bin

# Добавляем bin в текущей директории в PATH при запуске protoc
PROTOC = PATH="$$PATH:$(LOCAL_BIN)" protoc

LOMS_PROTO_PATH:="api/loms/v1"


PHONY: protoc-generate
protoc-generate: #.bin-deps .vendor-proto
	mkdir -p api/openapiv2
	protoc \
	-I ../loms/${LOMS_PROTO_PATH} \
	-I ".\..\vendor-proto" \
	--plugin=protoc-gen-go=$(LOCAL_BIN)/protoc-gen-go.exe \
	--go_out pkg/${LOMS_PROTO_PATH} \
	--go_opt paths=source_relative \
	--plugin=protoc-gen-go-grpc=$(LOCAL_BIN)/protoc-gen-go-grpc.exe \
	--go-grpc_out pkg/${LOMS_PROTO_PATH} \
	--go-grpc_opt paths=source_relative \
	../loms/${LOMS_PROTO_PATH}/loms.proto
	go mod tidy