package connection

import (
	"errors"
	"net"
	"time"
)

type UDP struct {
	address string
	options options
}

func DialUDP(address string, opts ...Option) *UDP {
	conn := &UDP{
		address: address,
		options: options{
			network: "udp",
		},
	}

	for _, opt := range opts {
		opt.apply(&conn.options)
	}

	return conn
}

func (c *UDP) connect() (*net.UDPConn, error) {
	conn, err := net.Dial(c.options.network, c.address)
	if err != nil {
		return nil, err
	}

	udp, ok := conn.(*net.UDPConn)
	if !ok {
		return nil, errors.New("failed cast connect to *net.UDPConn")
	}

	return udp, err
}

func (c *UDP) Read(b []byte) (n int, err error) {
	conn, err := c.connect()
	if err != nil {
		return -1, err
	}

	if c.options.readTimeout > 0 {
		if err = conn.SetReadDeadline(time.Now().Add(c.options.readTimeout)); err != nil {
			return -1., err
		}
	}

	return conn.Read(b)
}

func (c *UDP) Write(b []byte) (n int, err error) {
	conn, err := c.connect()
	if err != nil {
		return -1, err
	}

	if c.options.writeTimeout > 0 {
		if err = conn.SetWriteDeadline(time.Now().Add(c.options.readTimeout)); err != nil {
			return -1., err
		}
	}

	return conn.Write(b)
}

func (c *UDP) Close() error {
	return nil
}

func (c *UDP) Invoke(request []byte) (response []byte, err error) {
	conn, err := c.connect()
	if err != nil {
		return nil, err
	}

	if _, err = conn.Write(request); err != nil {
		return nil, err
	}

	b := make([]byte, bufferSize)

	n, _, err := conn.ReadFromUDP(b)
	if n > 0 {
		return b[:n], nil
	}

	return nil, nil
}
