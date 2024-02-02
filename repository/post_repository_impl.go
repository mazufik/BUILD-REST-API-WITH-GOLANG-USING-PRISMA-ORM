package repository

import (
	"context"
	"fmt"

	"github.com/mazufik/BUILD-REST-API-WITH-GOLANG-USING-PRISMA-ORM/helper"
	"github.com/mazufik/BUILD-REST-API-WITH-GOLANG-USING-PRISMA-ORM/model"
	"github.com/mazufik/BUILD-REST-API-WITH-GOLANG-USING-PRISMA-ORM/prisma/db"
)

type PostRepositoryImpl struct {
	Db *db.PrismaClient
}

func NewPostRepository(Db *db.PrismaClient) PostRepository {
	return &PostRepositoryImpl{Db: Db}
}

// Delete implements PostRepository.
func (p *PostRepositoryImpl) Delete(ctx context.Context, postId string) {
	result, err := p.Db.Post.FindUnique(db.Post.ID.Equals(postId)).Delete().Exec(ctx)
	helper.ErrorPanic(err)
	fmt.Println("Rows affected: ", result)
}

// FindAll implements PostRepository.
func (*PostRepositoryImpl) FindAll(ctx context.Context) []model.Post {
	panic("unimplemented")
}

// FindById implements PostRepository.
func (*PostRepositoryImpl) FindById(ctx context.Context, postId string) (model.Post, error) {
	panic("unimplemented")
}

// Save implements PostRepository.
func (*PostRepositoryImpl) Save(ctx context.Context, post model.Post) {
	panic("unimplemented")
}

// Update implements PostRepository.
func (*PostRepositoryImpl) Update(ctx context.Context, post model.Post) {
	panic("unimplemented")
}
