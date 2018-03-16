#[ProtoActors](http://proto.actor/) PoC

If proto files are modified the generation of .pb files must  [Protocol Buffers for Go with Gadgets](https://github.com/gogo/protobuf)

To compile messages.proto you must run in chatprotoactor/messages next command:
`protoc -I:.  -I=$GOPATH/src --gogoslick_out=plugins=grpc:. messages.proto `
 
To run Poc you must first run the server:

 `go run server/server.go`
 
 Then run one or more chat clients using:
 
`go run client/client.go`

Supported commands in client are:
- `quit`: to exit 
- `list`: to list connected clients
- `XXX <= hello`: To send a message to XXX 
- `hello`: To send a messaget to everyone

 
