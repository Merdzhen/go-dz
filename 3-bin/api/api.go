package api

type Client struct {
	Key string
}

func NewClient(key string) *Client {
	return &Client{Key: key}
}
