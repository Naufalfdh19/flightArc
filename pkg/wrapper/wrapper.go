package wrapper

type ApiResponse struct {
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
	Error   interface{} `json:"error,omitempty"`
}

func Response(data, err interface{}, message string) (ApiResponse) {
	return ApiResponse{
		Data: data,
		Message: message,
		Error: err,
	}
}
