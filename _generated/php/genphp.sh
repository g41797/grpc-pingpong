cd /home/g41797/Devs/try/proto
   
protoc \
   --php_out=../Pb      \
   --grpc_out=../Pb    \
   --plugin=protoc-gen-grpc=/home/g41797/Devs/grpc/cmake/build/grpc_php_plugin \
   *.proto
   

#   --proto_path=/home/g41797/Devs/try \