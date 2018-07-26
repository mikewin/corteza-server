package incoming

type (
	MessageCreate struct {
		ChannelID string `json:"cid"`
		Message   string `json:"msg"`
	}

	MessageUpdate struct {
		ID      string `json:"id"`
		Message string `json:"msg"`
	}

	MessageDelete struct {
		ID    string `json:"id"`
	}
)