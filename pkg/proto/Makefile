all: vierkantle.pb.go
.PHONY : all

vierkantle.pb.go: vierkantle.proto Makefile node_modules
	./node_modules/.bin/grpc_tools_node_protoc \
		--plugin=protoc-gen-ts_proto=./node_modules/.bin/protoc-gen-ts_proto \
		--ts_proto_out=../../cmd/ui/src/services \
		--ts_proto_opt=env=browser,outputServices=nice-grpc,outputServices=generic-definitions,outputJsonMethods=false,useExactTypes=false,esModuleInterop=true \
		--go_opt=paths=source_relative \
		--go_out=. \
		--go-grpc_opt=paths=source_relative \
		--go-grpc_opt=require_unimplemented_servers=false \
		--go-grpc_out=. \
		./vierkantle.proto

node_modules: yarn.lock
	yarn install --frozen-lockfile
	touch node_modules
