package server

import (
	"context"
	"net"

	"github.com/nilspolek/grpcCalculator/calculatorpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	calculatorpb.UnimplementedCalculatorServer
}

func (s *Server) Add(ctx context.Context, in *calculatorpb.CalculationRequest) (*calculatorpb.CalculationResponse, error) {
	result := in.GetLhs() + in.GetRhs()
	return &calculatorpb.CalculationResponse{Result: result}, nil
}

func (s *Server) Sub(ctx context.Context, in *calculatorpb.CalculationRequest) (*calculatorpb.CalculationResponse, error) {
	result := in.GetLhs() - in.GetRhs()
	return &calculatorpb.CalculationResponse{Result: result}, nil
}

func (s *Server) Mul(ctx context.Context, in *calculatorpb.CalculationRequest) (*calculatorpb.CalculationResponse, error) {
	result := in.GetLhs() * in.GetRhs()
	return &calculatorpb.CalculationResponse{Result: result}, nil
}

func (s *Server) Div(ctx context.Context, in *calculatorpb.CalculationRequest) (*calculatorpb.CalculationResponse, error) {
	result := in.GetLhs() / in.GetRhs()
	return &calculatorpb.CalculationResponse{Result: result}, nil
}

func (s *Server) Start() error {
	lis, err := net.Listen("tcp", ":12345")
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	calculatorpb.RegisterCalculatorServer(grpcServer, s)

	if err := grpcServer.Serve(lis); err != nil {
		return err
	}
	return nil
}
