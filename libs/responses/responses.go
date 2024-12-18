package responses

type Status struct {
	Status string `json:"status"`
}
type Message struct {
	Status
	Message string `json:"message"`
}

type Data struct {
	Status string
	Data   any `json:"data"`
}

type MessageData struct {
	Status
	Message string
	Data    any
}

type Error struct {
	Status string
	Errors any `json:"errors"`
}
