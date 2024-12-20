genproto:
	@protoc \
		--go_out=./services/users \
		--go-grpc_out=./services/users \
		services/users/user.proto