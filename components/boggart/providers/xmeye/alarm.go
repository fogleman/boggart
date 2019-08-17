package xmeye

import (
	"time"

	"github.com/kihamo/boggart/components/boggart/providers/xmeye/internal"
)

type AlertStreaming struct {
	client *Client
	alerts chan *internal.AlarmInfo
	errors chan error
	done   chan struct{}
}

func (c *Client) AlarmStart() error {
	_, err := c.Call(CmdGuardRequest, nil)
	return err
}

func (c *Client) AlarmStop() error {
	_, err := c.Call(CmdUnGuardRequest, nil)
	return err
}

func (c *Client) AlarmInfo() (*internal.AlarmInfo, error) {
	response := &internal.AlarmInfo{}

	err := c.CallWithResult(CmdAlarmRequest, nil, response)

	if err != nil {
		return nil, err
	}

	if response.AlarmInfo.Event == "" {
		return nil, nil
	}

	return response, nil
}

func (c *Client) AlarmStreaming() *AlertStreaming {
	s := &AlertStreaming{
		client: c,
		done:   make(chan struct{}),
		alerts: make(chan *internal.AlarmInfo),
		errors: make(chan error),
	}
	go s.start()

	return s
}

func (s *AlertStreaming) Close() error {
	close(s.done)
	return nil
}

func (s *AlertStreaming) NextAlarm() <-chan *internal.AlarmInfo {
	return s.alerts
}

func (s *AlertStreaming) NextError() <-chan error {
	return s.errors
}

func (s *AlertStreaming) start() {
	ticker := time.NewTicker(time.Second)

	defer func() {
		ticker.Stop()
		close(s.alerts)
		close(s.errors)
	}()

	var (
		alarmEnable bool
		err         error
	)

	for {
		select {
		case <-s.done:
			return

		case <-s.client.done:
			return

		case <-ticker.C:
			if !alarmEnable {
				if err = s.client.AlarmStart(); err != nil {
					s.errors <- err
					continue
				}

				alarmEnable = true
			}

			alert, err := s.client.AlarmInfo()
			if err != nil {
				s.errors <- err
			} else if alert != nil {
				s.alerts <- alert
			}
		}
	}
}
