package repository

import (
	belajargolang "belajar-golang-dasar"
	"belajar-golang-dasar/entity"
	"context"
	"fmt"
	"testing"
)


func TestCommentInsert(t *testing.T) {
	commentRepository := NewCommentRepository(belajargolang.GetConnection())

	ctx := context.Background()

	comment := entity.Comment{
		Email: "repository@test.com",
		Comment: "Test Repository",
	}

	result, err := commentRepository.Insert(ctx, comment)

	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestFindById(t *testing.T) {
	commentRepository := NewCommentRepository(belajargolang.GetConnection())

	ctx := context.Background()

	comment, err := commentRepository.FindById(ctx, 37)

	if err != nil {
		panic(err)
	}

	fmt.Println(comment)
}

func TestFindAll(t *testing.T) {
	commentRepository := NewCommentRepository(belajargolang.GetConnection())

	ctx := context.Background()

	comments, err := commentRepository.FindByAll(ctx)

	if err != nil {
		panic(err)
	}

	for _, comment := range comments {
		fmt.Println(comment)
	}

}