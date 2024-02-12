#!/bin/bash

rm -rf pb/*.go
protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
  --go-grpc_out=pb --go-grpc_opt=paths=source_relative --go-grpc_opt=require_unimplemented_servers=false \
  proto/*.proto