// Code generated by "enumer -type=DefaultAction -linecomment -text"; DO NOT EDIT.

package nethelpers

import (
	"fmt"
)

const _DefaultActionName = "acceptblock"

var _DefaultActionIndex = [...]uint8{0, 6, 11}

func (i DefaultAction) String() string {
	if i < 0 || i >= DefaultAction(len(_DefaultActionIndex)-1) {
		return fmt.Sprintf("DefaultAction(%d)", i)
	}
	return _DefaultActionName[_DefaultActionIndex[i]:_DefaultActionIndex[i+1]]
}

var _DefaultActionValues = []DefaultAction{0, 1}

var _DefaultActionNameToValueMap = map[string]DefaultAction{
	_DefaultActionName[0:6]:  0,
	_DefaultActionName[6:11]: 1,
}

// DefaultActionString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func DefaultActionString(s string) (DefaultAction, error) {
	if val, ok := _DefaultActionNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to DefaultAction values", s)
}

// DefaultActionValues returns all values of the enum
func DefaultActionValues() []DefaultAction {
	return _DefaultActionValues
}

// IsADefaultAction returns "true" if the value is listed in the enum definition. "false" otherwise
func (i DefaultAction) IsADefaultAction() bool {
	for _, v := range _DefaultActionValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalText implements the encoding.TextMarshaler interface for DefaultAction
func (i DefaultAction) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for DefaultAction
func (i *DefaultAction) UnmarshalText(text []byte) error {
	var err error
	*i, err = DefaultActionString(string(text))
	return err
}
