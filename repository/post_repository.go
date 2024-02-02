package repository

import (
	"context"

	"github.com/mazufik/BUILD-REST-API-WITH-GOLANG-USING-PRISMA-ORM/model"
)

type PostRepository interface {
	Save(ctx context.Context, post model.Post)
	Update(ctx context.Context, post model.Post)
	Delete(ctx context.Context, postId string)
	FindById(ctx context.Context, postId string) (model.Post, error)
	FindAll(ctx context.Context) []model.Post
}
