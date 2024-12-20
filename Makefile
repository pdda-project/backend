genproto:
# User proto
	@protoc \
		--go_out=./services/users \
		--go-grpc_out=./services/users \
		services/users/user.proto

# auth proto
	@protoc \
		--go_out=./services/auth \
		--go-grpc_out=./services/auth \
		services/auth/auth.proto

# Dataset proto
	@protoc \
		--go_out=./services/dataset \
		--go-grpc_out=./services/dataset \
		services/dataset/dataset.proto