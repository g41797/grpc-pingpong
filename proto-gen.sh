#!/bin/bash

rm -rf internal/pb/*.go
protoc --proto_path=proto --go_out=internal/pb --go_opt=paths=source_relative \
  --go-grpc_out=internal/pb --go-grpc_opt=paths=source_relative --go-grpc_opt=require_unimplemented_servers=false \
  proto/*.proto