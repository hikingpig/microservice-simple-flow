package cmd

import (
	"context"
	"database/sql"
	"flag"
	"fmt"

	// mysql driver
	_ "github.com/go-sql-driver/mysql"

	"github.com/hikingpig/microservice-simple-flow/pkg/protocol/grpc"
	"github.com/hikingpig/microservice-simple-flow/pkg/service/v1"
)

// Config is configuration for Server
type Config struct {
	GRPCPort string
	DatastoreDBHost string
	DatastoreDBUser string
	DatastoreDBPassword string
	DatastoreDBSchema string
}

func RunServer() error {
	ctx := context.Background()

	var cfg Config
	flag.StringVar(&cfg.GRPCPort, "grpc-port", "", "gRPC port to bind")
	flag.StringVar(&cfg.DatastoreDBHost, "db-host", "", "Database host")
	flag.StringVar(&cfg.DatastoreDBUser, "db-user", "", "Database user")
	flag.StringVar(&cfg.DatastoreDBPassword, "db-password", "", "Database password")
	flag.StringVar(&cfg.DatastoreDBSchema, "db-schema", "", "Database schema")
	flag.Parse()


	param := "parseTime=true"

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?%s",
		cfg.DatastoreDBUser,
		cfg.DatastoreDBPassword,
		cfg.DatastoreDBHost,
		cfg.DatastoreDBSchema,
		param)
	db, _ := sql.Open("mysql", dsn)

	defer db.Close()

	v1API := v1.NewToDoServiceServer(db)
	// now we use the service clearly with all the methods we declare in the service file
	// so the server is only interface, that has some method
	// how the method will work, is defined in the service
	return grpc.RunServer(ctx, v1API, cfg.GRPCPort)
}
