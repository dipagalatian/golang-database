package repository

import (
	"context"
	"fmt"
	golang_database "golang-database"
	"golang-database/entity"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestCommentInsert(t *testing.T) {
	
	commentRepository := NewCommentRepositoy(golang_database.GetConnectionDB())

	ctx := context.Background()
	comment := entity.Comment{
		Email: "repository@test.com",
		Comment: "Test Insert Comment from Repository",
	}

	result, err := commentRepository.Insert(ctx, comment)
	if err != nil {
		panic(err)
	}

	fmt.Println("Result TestCommentInsert:", result)
	
}  