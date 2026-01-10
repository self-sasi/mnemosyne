package api

type Client struct {
}

func NewClient() *Client {
	return &Client{}
}

func (client *Client) RegisterEngine(engine Engine) error {
	return nil
}

func (client *Client) UseRegistry() {
}
