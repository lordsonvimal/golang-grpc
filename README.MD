GRPC is a high performance RPC framework.

This Repository demonstrates different GRPC features using golang.

## Features
1. Data Serialization - GRPC uses protocol buffers as its serialization tool
2. HTTP - GRPC uses HTTP/2 by default
3. Works across Language and Platform
4. Streaming with integrated authentication

### Different Modes supported in GRPC
1. Unary
2. Server streaming
3. Client streaming
4. Bidirectional streaming

### Manual code generation

`protoc -Igreet/proto --go_out=. --go_opt=module=github.com/lordsonvimal/golang-grpc --go-grpc_out=. --go-grpc_opt=module=github.com/lordsonvimal/golang-grpc greet/proto/dummy.proto`

### Automatic code generation

`make greet`


### Workflow
1. Create proto file inside `<project>` directory
2. Run `make <project>`. This will generate the code to be consumed by server and client
3. Create server file and client file
4. Run `make <project>` to generate the build
5. Run `./bin/<project>/server` in one terminal tab to start the server
6. Run `./bin/<project>/client` in another terminal to run the client code

### ENV

#### Export path to enable protoc-gen-go command
`export PATH="$PATH:$(go env GOPATH)/bin"`