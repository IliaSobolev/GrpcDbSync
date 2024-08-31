package DbSyncRepository

import dbSync "grpcDbSync/proto"

type DbSyncRepository interface {
	Create(post *dbSync.Post) (*dbSync.Null, error)
	Update(post *dbSync.Post) (*dbSync.Null, error)
	Delete(post *dbSync.PostId) (*dbSync.Null, error)
	Get(post *dbSync.PostId) (*dbSync.Post, error)
}
