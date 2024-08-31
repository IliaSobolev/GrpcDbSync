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
	db  *pgxpool.Pool
	ctx context.Context
}

func New(db *pgxpool.Pool) *Repository {
	return &Repository{
		db:  db,
		ctx: context.Background(),
	}
}

func (h *Repository) Create(post *dbSync.Post) (*dbSync.Null, error) {
	_, err := h.db.Exec(h.ctx, createRequest, post.PostName)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (h *Repository) Update(post *dbSync.Post) (*dbSync.Null, error) {
	_, err := h.db.Exec(h.ctx, updateRequest, post.PostName, post.PostId.Id)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (h *Repository) Delete(post *dbSync.PostId) (*dbSync.Null, error) {
	_, err := h.db.Exec(h.ctx, deleteRequest, post.GetId())
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (h *Repository) Get(post *dbSync.PostId) (*dbSync.Post, error) {
	var (
		postId   int64
		postName string
		likes    int64
	)
	_ = h.db.QueryRow(h.ctx, getRequest, post.GetId()).Scan(&postId, &postName, &likes)
	if postId == 0 || postName == "" || likes == 0 {
		return nil, fmt.Errorf("no user was not found")
	}
	return &dbSync.Post{PostId: &dbSync.PostId{Id: postId}, PostName: postName, Likes: likes}, nil
}
