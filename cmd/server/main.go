package main

import (
	"github.com/hikingpig/microservice-simple-flow/pkg/cmd"
)

func main() {
	cmd.RunServer()
}

// run command
// server -grpc-port=9090 -db-host=localhost:3306 -db-user=root -db-password=1 -db-schema=mydb
