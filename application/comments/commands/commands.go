package comment_commands

import (
	contracts "rocky.my.id/git/mygram/application/comments/contracts"
)

type CommentCommands struct {
	Repository contracts.CommentRepositoryContract
}

func NewCommentCommands(repository contracts.CommentRepositoryContract) *CommentCommands {
	return &CommentCommands{Repository: repository}
}
