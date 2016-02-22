package irc

// Driver defines the IRC driver.
type Driver struct {
	host     string
	port     int
	username string
	channels []string
}

// New creates a new IRC driver.
func New(
	host string,
	port int,
	username string,
	channels []string) (*Driver, error) {
	driver := &Driver{
		host:     host,
		port:     port,
		username: username,
		channels: channels,
	}

	return driver, nil
}

func (d *Driver) Connect() {
}
