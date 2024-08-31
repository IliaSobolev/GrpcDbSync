package DbSyncRepository

import (
	"context"
	dbSync "grpcDbSync/proto"
)

type DbSyncRepository interface {
	Create(ctx context.Context, post *dbSync.Post) (*dbSync.Null, error)
	Update(ctx context.Context, post *dbSync.Post) (*dbSync.Null, error)
	Delete(ctx context.Context, post *dbSync.PostId) (*dbSync.Null, error)
	Get(ctx context.Context, post *dbSync.PostId) (*dbSync.Post, error)
}
