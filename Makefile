
.PHONY: generate
generate:
	go generate ./...
	KITEX_TOOL_USE_PROTOC=1 go tool kitex -module=mykitex -type="protobuf" -gen-path=./kitex_gen/protobuf -I ./proto ./proto/orders.proto
	KITEX_TOOL_USE_PROTOC=0 go tool kitex -module=mykitex -type="protobuf" -gen-path=./kitex_gen/prutal -I ./proto ./proto/orders.proto


benchmark:
	go test -benchmem -run=^$ -coverprofile=/tmp/vscode-godtJXfF/go-code-cover -bench . mykitex