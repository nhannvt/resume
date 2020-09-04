package auth

// Client represents a client uses sforum API.
type Client struct {

	// Name describe client name.
	name string

	// hashedKey is a hexadecimal hashed value of API key for the client.
	hashedKey string

	// role indicates the role of this client.
	// role limits the actions this client does.
	role *Role

	// Stages where client is available
	stages []string
}

func NewClient(name string, hashedKey string, role *Role, stages []string) *Client {
	return &Client{name, hashedKey, role, stages}
}

// Name returns Client's name
func (c *Client) Name() string {
	return c.name
}

// HashedKey returns Client's API Key hased value of a hexadecimal
func (c *Client) HashedKey() string {
	return c.hashedKey
}

// Role returns Client's role.
func (c *Client) Role() *Role {
	return c.role
}

// Stages return a list of stage name which client supports
func (c *Client) Stages() []string {
	return c.stages
}
