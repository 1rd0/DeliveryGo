gen:
	protoc \
      --proto_path=proto \
      --go_out=pkg/api \
      --go_opt=paths=source_relative \
      --go-grpc_out=pkg/api \
      --go-grpc_opt=paths=source_relative \
      proto/common/address.proto \
      proto/common/order.proto \
      proto/client.proto \
      proto/courier.proto \
      proto/order.proto

