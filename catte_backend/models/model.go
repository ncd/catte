package models

type Command struct {
	Action string `json:"action"`
	Room   string `json:"room"`
	Id     string `json:"id"`
	Data   string `json:"data"`
}

type ResponseCommand struct {
	Action string `json:"action"`
	Data   string `json:"data"`
}

type PlayData struct {
	Id   string `json:"id"`
	Data string `json:"data"`
}

type PlayerInfo struct {
	Id       string  `json:"id" form:"-"`
	Username string  `json:"username" form:"username"`
	Amount   float32 `json:"amount" form:"amount"`
}

type Deck []string
