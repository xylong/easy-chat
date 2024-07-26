goctl rpc protoc apps/social/rpc/social.proto --go_out=./apps/social/rpc --go-grpc_out=./apps/social/rpc --zrpc_out=./apps/social/rpc -style go_zero

goctl api go -api apps/social/api/social.api -dir apps/social/api -style go_zero

goctl model mysql ddl -src="./deploy/sql/user.sql" -dir="./apps/user/models/" -c -style go_zero

goctl model mysql ddl -src="./deploy/sql/social.sql" -dir="./apps/social/socialModels/" -c