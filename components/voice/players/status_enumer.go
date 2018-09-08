// Code generated by "enumer -type=Status -trimprefix=Status -output=status_enumer.go -transform=snake"; DO NOT EDIT.

package players

import (
	"fmt"
)

const _StatusName = "unknownstoppedplayingpause"

var _StatusIndex = [...]uint8{0, 7, 14, 21, 26}

func (i Status) String() string {
	if i < 0 || i >= Status(len(_StatusIndex)-1) {
		return fmt.Sprintf("Status(%d)", i)
	}
	return _StatusName[_StatusIndex[i]:_StatusIndex[i+1]]
}

var _StatusValues = []Status{0, 1, 2, 3}

var _StatusNameToValueMap = map[string]Status{
	_StatusName[0:7]:   0,
	_StatusName[7:14]:  1,
	_StatusName[14:21]: 2,
	_StatusName[21:26]: 3,
}

// StatusString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func StatusString(s string) (Status, error) {
	if val, ok := _StatusNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to Status values", s)
}

// StatusValues returns all values of the enum
func StatusValues() []Status {
	return _StatusValues
}

// IsAStatus returns "true" if the value is listed in the enum definition. "false" otherwise
func (i Status) IsAStatus() bool {
	for _, v := range _StatusValues {
		if i == v {
			return true
		}
	}
	return false
}
