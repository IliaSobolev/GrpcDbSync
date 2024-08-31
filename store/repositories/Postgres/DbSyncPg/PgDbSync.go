package DbSyncPg

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	dbSync "grpcDbSync/proto"
)

const createRequest = "INSERT INTO post(post_name) VALUES($1)"
const updateRequest = "UPDATE post SET post_name = $1 WHERE post_id = $2"
const deleteRequest = "DELETE FROM post WHERE post_id = $1"
const getRequest = "SELECT * FROM post WHERE post_id = $1"

type Repository struct {
	db *pgxpool.Pool
}

func New(db *pgxpool.Pool) *Repository {
	return &Repository{
		db: db,
	}
}

func (h *Repository) Create(ctx context.Context, post *dbSync.Post) (*dbSync.Null, error) {
	_, err := h.db.Exec(ctx, createRequest, post.PostName)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (h *Repository) Update(ctx context.Context, post *dbSync.Post) (*dbSync.Null, error) {
	_, err := h.db.Exec(ctx, updateRequest, post.PostName, post.PostId.Id)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (h *Repository) Delete(ctx context.Context, post *dbSync.PostId) (*dbSync.Null, error) {
	_, err := h.db.Exec(ctx, deleteRequest, post.GetId())
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (h *Repository) Get(ctx context.Context, post *dbSync.PostId) (*dbSync.Post, error) {
	var (
		postId   int64
		postName string
		likes    int64
	)
	_ = h.db.QueryRow(ctx, getRequest, post.GetId()).Scan(&postId, &postName, &likes)
	if postId == 0 || postName == "" || likes == 0 {
		return nil, fmt.Errorf("no user was not found")
	}
	return &dbSync.Post{PostId: &dbSync.PostId{Id: postId}, PostName: postName, Likes: likes}, nil
}
