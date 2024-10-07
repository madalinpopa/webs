package http

type Client interface {
	Get(url string) (string, error)
	Post(url string, data string) (string, error)
}

type client struct{}

func (c *client) Get(url string) (string, error) {
	return "", nil
}

func (c *client) Post(url string, data string) (string, error) {
	return "", nil
}

func NewClient() Client {
	return &client{}
}
