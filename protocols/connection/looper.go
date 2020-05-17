package connection

type Looper interface {
	Conn

	Loop() (responses <-chan []byte, errors <-chan error, done chan<- struct{})
}

type looper struct {
	Conn
}

func NewLooper(conn Conn) Looper {
	if i, ok := conn.(Looper); ok {
		return i
	}

	return &looper{
		Conn: conn,
	}
}

func (l *looper) Loop() (<-chan []byte, <-chan error, chan<- struct{}) {
	response := make(chan []byte)
	errors := make(chan error)
	done := make(chan struct{}, 1)

	go func() {
		buf := make([]byte, bufferSize)

		for {
			select {
			case <-done:
				return

			default:
				n, err := l.Conn.Read(buf)

				if n > 0 {
					go func(d []byte) {
						response <- d
					}(buf[:n])
				}

				if err != nil {
					go func(e error) {
						errors <- e
					}(err)
				}
			}
		}
	}()

	return response, errors, done
}
