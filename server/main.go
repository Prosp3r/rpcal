package main

import (
	pb "github.com/Prosp3r/rpcal/proto"
)

const (
	port = ":50051"
)

type repository interface {
	Create(*pb.Calendar) ([]*pb.Calendar, error)
}

func main() {

}
