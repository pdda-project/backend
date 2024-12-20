package responses

type Response[T any] struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    T      `json:"data"`
}

type ResponseMessage struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type ResponseData[T any] struct {
	Status string `json:"status"`
	Data   T      `json:"data"`
}
