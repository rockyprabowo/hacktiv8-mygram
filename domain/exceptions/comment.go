package exceptions

import (
	"fmt"
)

var CommentNotFoundError = fmt.Errorf("comment %w", EntityNotFound)
