package responses

type InfoResult struct {
	Message string `json:"message"`
}

type ErrorResult struct {
	Error any `json:"error"`
	Code  int `json:"code,omitempty"`
}
