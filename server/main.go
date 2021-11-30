package main

import (
	"context"
	"sync"

	pb "github.com/Prosp3r/rpcal/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

type repository interface {
	Create(*pb.Calendar) ([]*pb.Calendar, error)
	GetAll() []*pb.Calendar
}

type Repository struct {
	mu sync.RWMutex
	calendars []*pb.Calendar
}

func (repo *Repository) Create(calendar *pb.Calendar) (*pb.Calendar, error){

	repo.mu.Lock()
	repo.calendars = append(repo.calendars, calendar)
	repo.mu.Unlock()
	return calendar, nil
}

type service struct {
	repo repository
}

func (s *service) CreateCalendar(ctx context.Context, req *pb.Calendar) (*pb.Response, error){
	calendar, err := s.repo.Create(req)
	if err != nil {
		return nil, err
	}
	return &pb.Response{Created: true, Calendar: calendar}, nil
}

func main() {

}
