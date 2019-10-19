package main

import (
	"context"
	"fmt"
	"log"
	"net"

	reservations2 "github.com/docker/envoy-grpc-transcoder/reservation"
	"github.com/golang/protobuf/ptypes/empty"
	uuid "github.com/satori/go.uuid"
	"google.golang.org/grpc"
)

type Server struct{}

var cache map[string]reservations2.Reservation

func (s Server) CreateReservation(ctx context.Context, req *reservations2.CreateReservationRequest) (*reservations2.Reservation, error) {
	var err error
	fmt.Print("CreateReservation: %v", req)
	r := reservations2.Reservation{
		Id:        uuid.Must(uuid.NewV4(), err).String(),
		Title:     req.Reservation.Title,
		Venue:     req.Reservation.Venue,
		Room:      req.Reservation.Room,
		Attendees: req.Reservation.Attendees,
	}
	cache[r.Id] = r

	return &r, nil
}

func (s Server) DeleteReservation(ctx context.Context, req *reservations2.DeleteReservationRequest) (*empty.Empty, error) {
	fmt.Printf("DeleteReservation: %v", req)
	delete(cache, req.Id)
	return &empty.Empty{}, nil
}

func (s Server) GetReservation(ctx context.Context, req *reservations2.GetReservationRequest) (*reservations2.Reservation, error) {
	fmt.Printf("GetReservation: %s", req)
	if v, ok := cache[req.Id]; ok {
		return &v, nil
	}
	return nil, nil
}

func (s Server) ListReservations(req *reservations2.ListReservationRequest, resp reservations2.ReservationService_ListReservationsServer) error {
	fmt.Printf("ListReservations: %v", req)
	return nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 53000))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	reservations2.RegisterReservationServiceServer(grpcServer, &Server{})

	cache = make(map[string]reservations2.Reservation)

	fmt.Println("Starting GRPC server")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
