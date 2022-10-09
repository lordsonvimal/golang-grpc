Commands

Manual code generation

protoc -Igreet/proto --go_out=. --go_opt=module=github.com/lordsonvimal/golang-grpc --go-grpc_out=. --go-grpc_opt=module=github.com/lordsonvimal/golang-grpc greet/proto/dummy.proto

For automatic generation

make greet

Export path to enable protoc-gen-go command
export PATH="$PATH:$(go env GOPATH)/bin"