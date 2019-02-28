// set up protoc:
	// binaries => bin
	// mkidir third_party
	// copy all files in include to third_party

// mkdir cmd/client-grpc
	// create file main.go

	req1 := v1.CreateRequest{
		Api: apiVersion,
		ToDo: &v1.ToDo{
			Title:       "title (" + pfx + ")",
			Description: "description (" + pfx + ")",
			Reminder:    reminder,
		},
	}
	res1, err := c.Create(ctx, &req1)

// prepare go protobuf to receive message
	// mkdir -p api/proto/v1
	// create file todo-service.proto
			rpc Create(CreateRequest) returns (CreateResponse){
      option (google.api.http) = {
        post: "/v1/todo"
        body: "*"
				};
			}
	// create a folder for the go protobuf
		// mkdir -p pkg/api/v1
		// protoc --proto_path=api/proto/v1 --proto_path=third_party --go_out=plugins=grpc:pkg/api/v1 todo-service.proto

	// create a service to handle request
		// mkdir -p pkg/service/v1/
		// create a file todo-service.go
				func (s *toDoServiceServer) Create(ctx context.Context, req *v1.CreateRequest) (*v1.CreateResponse, error) {
					c, err := s.connect(ctx)
					defer c.Close()
					reminder, err := ptypes.Timestamp(req.ToDo.Reminder)
					res, err := c.ExecContext(ctx, "INSERT INTO ToDo(`Title`, `Description`, `Reminder`) VALUES(?, ?, ?)",
						req.ToDo.Title, req.ToDo.Description, reminder)
					id, err := res.LastInsertId()
					return &v1.CreateResponse{
						Api: apiVersion,
						Id:  id,
					}, nil
				}
		// create a grpc server:
			// mkdir pkg/protocol/grpc
					listen, err := net.Listen("tcp", ":"+port)
					server := grpc.NewServer()
					v1.RegisterToDoServiceServer(server, v1API)
					return server.Serve(listen)
		// create a go server:
			// mkdr pkg/cmd/server
			// cmd will call protocol in pkg
			// create a file server.go wit RunServer function
		// create a main cmd server that can run the pkg/cmd go server
			// mkdir cmd
			// create file server.go