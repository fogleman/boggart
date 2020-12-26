// Code generated by "enumer -type=DrumStatus -trimprefix=DrumStatus -output=drum_status_enumer.go"; DO NOT EDIT.

//
package pantum

import (
	"fmt"
)

const _DrumStatusName = "UnknownNormalUninstalledUnmatchedInvalidExpired"

var _DrumStatusIndex = [...]uint8{0, 7, 13, 24, 33, 40, 47}

func (i DrumStatus) String() string {
	i -= -1
	if i < 0 || i >= DrumStatus(len(_DrumStatusIndex)-1) {
		return fmt.Sprintf("DrumStatus(%d)", i+-1)
	}
	return _DrumStatusName[_DrumStatusIndex[i]:_DrumStatusIndex[i+1]]
}

var _DrumStatusValues = []DrumStatus{-1, 0, 1, 2, 3, 4}

var _DrumStatusNameToValueMap = map[string]DrumStatus{
	_DrumStatusName[0:7]:   -1,
	_DrumStatusName[7:13]:  0,
	_DrumStatusName[13:24]: 1,
	_DrumStatusName[24:33]: 2,
	_DrumStatusName[33:40]: 3,
	_DrumStatusName[40:47]: 4,
}

// DrumStatusString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func DrumStatusString(s string) (DrumStatus, error) {
	if val, ok := _DrumStatusNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to DrumStatus values", s)
}

// DrumStatusValues returns all values of the enum
func DrumStatusValues() []DrumStatus {
	return _DrumStatusValues
}

// IsADrumStatus returns "true" if the value is listed in the enum definition. "false" otherwise
func (i DrumStatus) IsADrumStatus() bool {
	for _, v := range _DrumStatusValues {
		if i == v {
			return true
		}
	}
	return false
}
