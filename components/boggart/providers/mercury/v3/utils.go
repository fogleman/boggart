package v3

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

// Меркурий 230, 231, 233, 234, 236
func ConvertSerialNumber(serial string) byte {
	// Отделите три последние цифры серийного номера, это будет число N.
	number, _ := strconv.ParseInt(serial[len(serial)-3:], 10, 0)

	// Если N>=240 адресом являются две последние цифры серийного номера.
	if number >= 240 {
		number, _ = strconv.ParseInt(serial[len(serial)-2:], 10, 0)
		return byte(number)
	}

	// Если N<240 адресом являются три последние цифры.
	if number < 240 {
		return byte(number)
	}

	return 1
}

func ResponseError(response *Response) error {
	return PayloadError(response.Payload)
}

func PayloadError(payload []byte) error {
	switch payload[0] {
	case ResponseCodeBadRequest:
		return errors.New("bad request")
	case ResponseCodeInternalError:
		return errors.New("internal error")
	case ResponseCodeAccessDenied:
		return errors.New("access denied")
	case ResponseCodeTimeCorrectFiled:
		return errors.New("correct time failed")
	case ResponseCodeChannelClosed:
		return errors.New("channel is closed")
	}

	return nil
}

func ParseFirmwareVersion(data []byte) string {
	return fmt.Sprintf("%d.%d.%d", data[0], data[1], data[2])
}

func ParseSerialNumber(data []byte) string {
	return fmt.Sprintf("%d%d%d%d", data[0], data[1], data[2], data[3])
}

func ParseBuildDate(data []byte) time.Time {
	return time.Date(2000+int(data[2]), time.Month(int(data[1])), int(data[0]), 0, 0, 0, 0, time.UTC)
}

// TODO:
func ParseType(data []byte) *Type {
	return nil
}
