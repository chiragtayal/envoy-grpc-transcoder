# Transcoding gRPC to HTTP/JSON

Showing how to expose a gRPC service as a HTTP/JSON api.


# Getting Started

1. Clone these repositories in your GOPATH directory
	`git clone git@github.com:chiragtayal/envoy-grpc-transcoder.git`
	
	`git clone https://github.com/googleapis/googleapis` 
	This is used for annotations in proto file

2. Build descriptor file and pn.go file using following commands
	`export GOOGLEAPIS_DIR=/Users/chiragtayal/docker/grpc/src/github.com/googleapis`

	`protoc -I$GOOGLEAPIS_DIR -I. --include_imports --include_source_info   --descriptor_set_out=proto.pb reservation/reservation.proto` 
	
	`protoc -I/usr/local/protoc/include -I. -I$GOPATH/src -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --go_out=plugins=grpc:. reservation.proto`

3. Start GRPC server 
	`./main &`

4. Start Envoy docker container
	`docker run -it --rm --name envoy --network host -v "$(pwd)/proto.pb:/data/proto.pb:ro" -v "$(pwd)/envoy.yml:/etc/envoy/envoy.yaml:ro" envoyproxy/envoy`


# Test the deployment

1. curl -v -XPOST http://localhost:51051/reservations -H 'Content-Type: application/json' -d '{"title": "Lunchmeeting2", "venue": "JDriven Coltbaan 3", "room": "atrium", "timestamp": "2018-10-10T11:12:13", "attendees": [{"firstName": "Jimmy","lastName": "Jones"},{"firstName": "Dennis", "lastName": "Richie"}]}'

2. curl -v http://localhost:51051/reservations/<ID> 
