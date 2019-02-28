package v1

import (
	"context"
	"database/sql"

	"github.com/golang/protobuf/ptypes"

	"github.com/hikingpig/microservice-simple-flow/pkg/api/v1"
)

const (
	apiVersion = "v1"
)

type toDoServiceServer struct {
	db *sql.DB
}

func NewToDoServiceServer(db *sql.DB) v1.ToDoServiceServer {
	return &toDoServiceServer{db: db}
}

func (s *toDoServiceServer) connect(ctx context.Context) (*sql.Conn, error) {
	c, _ := s.db.Conn(ctx)
	return c, nil
}

func (s *toDoServiceServer) Create(ctx context.Context, req *v1.CreateRequest) (*v1.CreateResponse, error) {
	c, _ := s.connect(ctx)

	defer c.Close()

	reminder, _ := ptypes.Timestamp(req.ToDo.Reminder)
	res, _ := c.ExecContext(ctx, "INSERT INTO ToDo(`Title`, `Description`, `Reminder`) VALUES(?, ?, ?)",
		req.ToDo.Title, req.ToDo.Description, reminder)
	id, _ := res.LastInsertId()

	return &v1.CreateResponse{
		Api: apiVersion,
		Id:  id,
	}, nil
}
