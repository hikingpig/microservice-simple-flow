  // 1. http client send request:
 	resp, err := http.Post(*address+"/v1/todo", "application/json", strings.NewReader(fmt.Sprintf(`
		{
			"api":"v1",
			"toDo": {
				"title":"title (%s)",
				"description":"description (%s)",
				"reminder":"%s"
			}
		}
	`, pfx, pfx, pfx)))

// ****** how can request body turn into message form?
// what happen with path param?

	// 1.2. grpc client send request:
	// v1 is compiled from proto file. it has the message types
	// conn is connection to grpc server
	// c is the client service created from v1 and grpc server
	req1 := v1.CreateRequest{
		Api: apiVersion,
		ToDo: &v1.ToDo{
			Title:       "title (" + pfx + ")",
			Description: "description (" + pfx + ")",
			Reminder:    reminder,
		},
	}
	res1, err := c.Create(ctx, &req1)

	// 2. prepare to receive that request:

		// protocol buffer/gateway
		rpc Create(CreateRequest) returns (CreateResponse){
      option (google.api.http) = {
        post: "/v1/todo"
        body: "*"
      };
    }

		// handle grpc request in go service
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

		// register service in grpc server
				listen, err := net.Listen("tcp", ":"+port)
				server := grpc.NewServer()
				v1.RegisterToDoServiceServer(server, v1API)
				// register method is from the proto.
				// v1API type is ToDoServiceServer
				// a service will have register server and register client
				// and also 2 structs, serviceServer and serviceClient
				return server.Serve(listen)

		// register service in http server
					mux := runtime.NewServeMux()
					opts := []grpc.DialOption{grpc.WithInsecure()}
					v1.RegisterToDoServiceHandlerFromEndpoint(ctx, mux, "localhost:"+grpcPort, opts)
					// register from end point will have the grpc handler
					// and the gateway for each path. so the http request will
					// be converted to grpc message and passed to grpc server
					// grpc response will be converted to http response and send
					// to the http server
					// so converting between the 2 data types is handled by the gateway

					// not figure out yet how this conversion happens correctly
					srv := &http.Server{
						Addr:    ":" + httpPort,
						Handler: mux,
					}
					return srv.ListenAndServe()
