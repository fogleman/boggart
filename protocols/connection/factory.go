package connection

import (
	"errors"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/kihamo/boggart/protocols/serial"
)

type option int64

const (
	OptionInvoker option = iota
	OptionDumper
)

func New(dsn string) (conn Conn, err error) {
	if dsn == "" {
		return nil, errors.New("DSN is empty")
	}

	u, err := url.Parse(dsn)
	if err != nil {
		return nil, err
	}

	switch u.Scheme {
	case "tcp", "tcp4", "tcp6":
		options := []Option{
			WithNetwork(u.Scheme),
		}

		for key, value := range u.Query() {
			switch strings.ToLower(key) {
			case "read-timeout":
				v, err := time.ParseDuration(value[0])
				if err != nil {
					return nil, err
				}

				options = append(options, WithReadTimeout(v))

			case "write-timeout":
				v, err := time.ParseDuration(value[0])
				if err != nil {
					return nil, err
				}

				options = append(options, WithWriteTimeout(v))

			case "once":
				v, err := strconv.ParseBool(value[0])
				if err != nil {
					return nil, err
				}

				options = append(options, WithOnce(v))
			}
		}

		conn = Dial(u.Host, options...)

	case "udp", "udp4", "udp6", "unixgram":
		options := []Option{
			WithNetwork(u.Scheme),
		}

		for key, value := range u.Query() {
			switch strings.ToLower(key) {
			case "read-timeout":
				v, err := time.ParseDuration(value[0])
				if err != nil {
					return nil, err
				}

				options = append(options, WithReadTimeout(v))

			case "write-timeout":
				v, err := time.ParseDuration(value[0])
				if err != nil {
					return nil, err
				}

				options = append(options, WithWriteTimeout(v))

			case "once":
				v, err := strconv.ParseBool(value[0])
				if err != nil {
					return nil, err
				}

				options = append(options, WithOnce(v))
			}
		}

		conn = Dial(u.Host, options...)

	case "serial":
		options := []serial.Option{
			serial.WithAddress(u.EscapedPath()),
		}

		for key, value := range u.Query() {
			switch strings.ToLower(key) {
			case "baudrate":
				v, err := strconv.ParseInt(value[0], 10, 64)
				if err != nil {
					return nil, err
				}

				options = append(options, serial.WithBaudRate(int(v)))

			case "databits":
				v, err := strconv.ParseInt(value[0], 10, 64)
				if err != nil {
					return nil, err
				}

				options = append(options, serial.WithDataBits(int(v)))

			case "stopbits":
				v, err := strconv.ParseInt(value[0], 10, 64)
				if err != nil {
					return nil, err
				}

				options = append(options, serial.WithStopBits(int(v)))

			case "parity":
				options = append(options, serial.WithParity(value[0]))

			case "timeout":
				v, err := time.ParseDuration(value[0])
				if err != nil {
					return nil, err
				}

				options = append(options, serial.WithTimeout(v))
			}
		}

		conn = NewIO(serial.Dial(options...))

	default:
		err = errors.New("unknown connection type for DSN " + dsn)
	}

	return
}

func NewWithOptions(dsn string, options ...option) (conn Conn, err error) {
	conn, err = New(dsn)

	if err == nil {
		for _, opt := range options {
			switch opt {
			case OptionInvoker:
				conn = NewInvoker(conn)
			case OptionDumper:
				conn = NewDumper(conn)
			}
		}
	}

	return conn, err
}