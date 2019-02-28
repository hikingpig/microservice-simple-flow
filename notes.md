set up binaries:
  protobuf => extract => add bin to path

mkdir third_party

copy all in the protoc/include to third_party folder

go get -u github.com/golang/protobuf/protoc-gen-go

mkdir -p api/proto/v1

write proto file



// pkg folder to store go protobuf file
mkdir -p pkg/api/v1

convert proto file to go protobuf
protoc --proto_path=api/proto/v1 --proto_path=third_party --go_out=plugins=grpc:pkg/api/v1 todo-service.proto

// create a service
mkdir -p pkg/service/v1/

create a service file todo-service.go

// create gRPC server
mkdir pkg/protocol/grpc

create file server.go

// create a go server
mkdir pkg/cmd/server

create file server.go that will include grpc server

// create a cmd go to run go server
mkdir cmd/server

create file main.go

// create grpc client
mkdir cmd/client-grpc

create file main.go




