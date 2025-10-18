package response

type Error struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}
