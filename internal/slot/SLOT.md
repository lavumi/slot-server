# Slot Server
grpc slot server for MSA


## Run protobuf Compiler
```bash
rm -rf api/proto/**
protoc --proto_path=proto --go_out=api --go_opt=paths=import --go-grpc_out=api --go-grpc_opt=paths=import proto/base/*.proto proto/*.proto
```