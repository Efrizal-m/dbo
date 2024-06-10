package interfaces

type SuccessResponseWrapper struct {
	// Example: {"data": "your-response-data"}
	Success string `json:"data"`
}

type ErrorResponseWrapper struct {
	// Example: {"error": "your-error-value"}
	Error error `json:"error"`
}
