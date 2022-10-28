package reply

type Relations struct {
	UserID int64  `json:"user_id"`
	State  string `json:"state"`
	Type   string `json:"type"`
}
