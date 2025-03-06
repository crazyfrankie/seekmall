.PHONY: gen-user
gen-user:
	@protoc --go_out=./rpc_gen --go-grpc_out=./rpc_gen --grpc-gateway_out=./rpc_gen ./idl/seekmall/user.proto

.PHONY: gen-sm
gen-sm:
	@protoc --go_out=./rpc_gen --go-grpc_out=./rpc_gen --grpc-gateway_out=./rpc_gen ./idl/seekmall/sm.proto

.PHONY: gen-product
gen-product:
	@protoc --go_out=./rpc_gen --go-grpc_out=./rpc_gen --grpc-gateway_out=./rpc_gen ./idl/seekmall/product.proto

.PHONY: gen-cart
gen-cart:
	@protoc --go_out=./rpc_gen --go-grpc_out=./rpc_gen --grpc-gateway_out=./rpc_gen ./idl/seekmall/cart.proto

.PHONY: gen-payment
gen-payment:
	@protoc --go_out=./rpc_gen --go-grpc_out=./rpc_gen --grpc-gateway_out=./rpc_gen ./idl/seekmall/payment.proto