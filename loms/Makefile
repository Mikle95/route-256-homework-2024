LOCAL_BIN:=$(CURDIR)/../bin

# Добавляем bin в текущей директории в PATH при запуске protoc
PROTOC = PATH="$$PATH:$(LOCAL_BIN)" protoc

NOTES_PROTO_PATH:="api/loms/v1"

PHONY: .protoc-generate
.protoc-generate: #.bin-deps .vendor-proto
	mkdir -p api/openapiv2
	protoc \
	-I ${NOTES_PROTO_PATH} \
	-I ".\..\vendor-proto" \
	--plugin=protoc-gen-go=$(LOCAL_BIN)/protoc-gen-go.exe \
	--go_out pkg/${NOTES_PROTO_PATH} \
	--go_opt paths=source_relative \
	--plugin=protoc-gen-go-grpc=$(LOCAL_BIN)/protoc-gen-go-grpc.exe \
	--go-grpc_out pkg/${NOTES_PROTO_PATH} \
	--go-grpc_opt paths=source_relative \
	--plugin=protoc-gen-validate=$(LOCAL_BIN)/protoc-gen-validate.exe \
	--validate_out="lang=go,paths=source_relative:pkg/${NOTES_PROTO_PATH}" \
	--plugin=protoc-gen-grpc-gateway=$(LOCAL_BIN)/protoc-gen-grpc-gateway.exe \
	--grpc-gateway_out pkg/${NOTES_PROTO_PATH} \
	--grpc-gateway_opt logtostderr=true --grpc-gateway_opt paths=source_relative --grpc-gateway_opt generate_unbound_methods=true \
	--plugin=protoc-gen-openapiv2=$(LOCAL_BIN)/protoc-gen-openapiv2.exe \
	--openapiv2_out api/openapiv2 \
	--openapiv2_opt logtostderr=true \
	${NOTES_PROTO_PATH}/loms.proto
	go mod tidy