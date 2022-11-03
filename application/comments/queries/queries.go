package comment_queries

import (
	contracts "rocky.my.id/git/mygram/application/comments/contracts"
)

type CommentQueries struct {
	Repository contracts.CommentRepositoryContract
}

func NewCommentQueries(repository contracts.CommentRepositoryContract) *CommentQueries {
	return &CommentQueries{Repository: repository}
}
