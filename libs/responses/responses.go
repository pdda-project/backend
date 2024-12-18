package responses

type Status struct {
	Status string `json:"status"`
}
type Message struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type Data struct {
	Status string `json:"status"`
	Data   any    `json:"data"`
}

type MessageData struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type Error struct {
	Status string `json:"status"`
	Errors any    `json:"errors"`
}
