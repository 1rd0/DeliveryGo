PROTO_DIR = proto
OUT_DIR = .

gen_order:
	protoc --proto_path=$(PROTO_DIR) \
	  --go_out=$(OUT_DIR) \
	  --go-grpc_out=$(OUT_DIR) \
	  $(PROTO_DIR)/order.proto
gen_user:
	protoc --proto_path=$(PROTO_DIR) \
	  --go_out=$(OUT_DIR) \
	  --go-grpc_out=$(OUT_DIR) \
	  $(PROTO_DIR)/user.proto
gen_all:
	make gen_order
	make gen_user