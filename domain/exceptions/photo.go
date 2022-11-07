package exceptions

import (
	"fmt"
)

var PhotoNotFoundError = fmt.Errorf("photo %w", EntityNotFound)
