package http_exceptions

import "errors"

var MalformedPayload = errors.New("malformed payload")
var MissingPayload = errors.New("missing payload")
