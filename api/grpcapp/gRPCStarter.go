package grpcapp

import (
	"fmt"
	"google.golang.org/grpc"
	"grpcDbSync/api/grpcapp/controllers/DbSyncController"
	"grpcDbSync/store/repositories/Postgres/DbSyncPg"
	"log/slog"
	"net"
)

type App struct {
	grpc *grpc.Server
	port int
}

func New(port int, repository *DbSyncPg.Repository) *App {
	grpcServer := grpc.NewServer()
	DbSyncController.Register(grpcServer, repository)
	return &App{
		grpc: grpcServer,
		port: port,
	}
}

func (a *App) Start() error {
	slog.Info("starting grpc server")
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		return err
	}
	if err := a.grpc.Serve(lis); err != nil {
		return err
	}
	return nil
}

func (a *App) Stop() {
	slog.Info("stopping grpc server")
	a.grpc.GracefulStop()
}
