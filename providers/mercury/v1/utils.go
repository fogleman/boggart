package v1

import (
	"encoding/hex"
	"fmt"
	"strconv"
	"time"
)

func ConvertSerialNumber(serial string) []byte {
	number, _ := strconv.ParseInt(serial[len(serial)-6:], 10, 0)
	h, _ := hex.DecodeString(fmt.Sprintf("%06x", number))

	return h
}

func ParseInt(data ...byte) int64 {
	result, _ := strconv.ParseInt(hex.EncodeToString(data), 10, 0)
	return result
}

func ParseUint(data ...byte) uint64 {
	return uint64(ParseInt(data...))
}

func ParseFloat(data ...byte) float64 {
	result, _ := strconv.ParseFloat(hex.EncodeToString(data), 10)
	return result
}

func ParseDate(data []byte, location *time.Location) time.Time {
	return time.Date(
		2000+int(ParseInt(data[2])),
		time.Month(ParseInt(data[1])),
		int(ParseInt(data[0])),
		0,
		0,
		0,
		0,
		location)
}

func ParseDatetime(data []byte, location *time.Location) time.Time {
	return time.Date(
		2000+int(ParseInt(data[6])),
		time.Month(ParseInt(data[5])),
		int(ParseInt(data[4])),
		int(ParseInt(data[1])),
		int(ParseInt(data[2])),
		int(ParseInt(data[3])),
		0,
		location)
}
