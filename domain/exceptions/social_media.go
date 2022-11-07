package exceptions

import (
	"fmt"
)

var SocialMediaNotFoundError = fmt.Errorf("social media %w", EntityNotFound)
