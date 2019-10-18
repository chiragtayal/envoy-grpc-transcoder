package main

import (
  "fmt"
  reservations2 "github.com/docker/reservation/reservation"
  "github.com/golang/protobuf/ptypes/empty"
  "log"
  "net"
  "google.golang.org/grpc"
  "context"
)

type Server struct {}

func (s Server) CreateReservation(ctx context.Context, req *reservations2.CreateReservationRequest) (*reservations2.Reservation, error) {
  fmt.Print("CreateReservation: %v", req)
  return &reservations2.Reservation{}, nil
}

func (s Server) DeleteReservation(ctx context.Context, req *reservations2.DeleteReservationRequest) (*empty.Empty, error) {
  fmt.Print("DeleteReservation: %v", req)
  return &empty.Empty{}, nil
}

func (s Server) GetReservation(ctx context.Context, req *reservations2.GetReservationRequest) (*reservations2.Reservation, error) {
  fmt.Print("GetReservation: %s", req)
  return &reservations2.Reservation{}, nil
}

func (s Server) ListReservations(req *reservations2.ListReservationRequest, resp reservations2.ReservationService_ListReservationsServer) error {
  fmt.Print("ListReservations: %v", req)
  return nil
}

func main() {
  lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 53000))
  if err != nil {
    log.Fatalf("failed to listen: %v", err)
  }

  grpcServer := grpc.NewServer()

  reservations2.RegisterReservationServiceServer(grpcServer, &Server{})

  fmt.Println("Starting GRPC server")
  if err := grpcServer.Serve(lis); err != nil {
    log.Fatalf("failed to serve: %s", err)
  }
}
