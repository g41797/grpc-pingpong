protoc \
   --php_out=../pb      \
   --grpc_out=../pb    \
   --plugin=protoc-gen-grpc=grpc_php_plugin \
   *.proto

