# Slot Server
grpc slot server for MSA



### Install
```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

## Run protobuf Compiler
```bash
rm -rf api/proto/**
protoc --proto_path=proto --go_out=api --go_opt=paths=import --go-grpc_out=api --go-grpc_opt=paths=import proto/base/*.proto proto/*.proto
```