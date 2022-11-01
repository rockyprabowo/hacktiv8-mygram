package comment_use_cases

import contracts "rocky.my.id/git/mygram/application/comments/contracts"

type CommentUseCases struct {
	Commands contracts.CommentCommandsContract
	Queries  contracts.CommentQueriesContract
}

func NewCommentUseCases(commands contracts.CommentCommandsContract, queries contracts.CommentQueriesContract) *CommentUseCases {
	return &CommentUseCases{Commands: commands, Queries: queries}
}
