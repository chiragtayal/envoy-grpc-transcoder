# envoy-grpc-transcoder

Build descriptor file
1. protoc -I$GOOGLEAPIS_DIR -I. --include_imports --include_source_info   --descriptor_set_out=proto.pb reservation/reservation.proto 
2. protoc -I/usr/local/protoc/include -I. -I$GOPATH/src -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --go_out=plugins=grpc:. reservation.proto

Run envoy in docker in host network

docker run -it --rm --name envoy --network host -v "$(pwd)/proto.pb:/data/proto.pb:ro" -v "$(pwd)/envoy.yml:/etc/envoy/envoy.yaml:ro" envoyproxy/envoy
