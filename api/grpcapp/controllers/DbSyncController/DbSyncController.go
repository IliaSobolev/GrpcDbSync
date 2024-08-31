package DbSyncController

import (
	"context"
	"google.golang.org/grpc"
	dbSync "grpcDbSync/proto"
	"grpcDbSync/store/repositories/Postgres/DbSyncPg"
	"log/slog"
)

type ServerApi struct {
	dbSync.UnimplementedDbSyncServer
	repo *DbSyncPg.Repository
}

func Register(grpc *grpc.Server, repo *DbSyncPg.Repository) {
	dbSync.RegisterDbSyncServer(grpc, &ServerApi{repo: repo})
}

func (s *ServerApi) Create(ctx context.Context, req *dbSync.Post) (*dbSync.Null, error) {
	slog.Info("Creating post with post_name: ", req.PostName)
	_, err := s.repo.Create(ctx, req)
	if err != nil {
		slog.Warn(err.Error())
		return nil, err
	}
	return &dbSync.Null{}, nil
}

func (s *ServerApi) Update(ctx context.Context, req *dbSync.Post) (*dbSync.Null, error) {
	slog.Info("Updating post with post_id: ", req.GetPostId())
	_, err := s.repo.Update(ctx, req)
	if err != nil {
		slog.Warn(err.Error())
		return nil, err
	}
	return &dbSync.Null{}, nil
}

func (s *ServerApi) Delete(ctx context.Context, req *dbSync.PostId) (*dbSync.Null, error) {
	slog.Info("Updating post with post_id: ", req.GetId())
	_, err := s.repo.Delete(ctx, req)
	if err != nil {
		slog.Warn(err.Error())
		return nil, err
	}
	return &dbSync.Null{}, nil
}

func (s *ServerApi) Get(ctx context.Context, req *dbSync.PostId) (*dbSync.Post, error) {
	slog.Info("Getting post with post_id: ", req.GetId())
	post, err := s.repo.Get(ctx, req)
	if err != nil {
		slog.Warn(err.Error())
		return nil, err
	}
	return post, nil
}
