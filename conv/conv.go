package conv

import (
	"errors"
)

var (
	ErrInterfaceToStringMap = errors.New("converting interface to map[string]string failed")
	ErrInterfaceToString    = errors.New("converting interface to string failed")
	ErrInterfaceToBool      = errors.New("converting interface to bool failed")
	ErrInterfaceToInt64     = errors.New("converting interface to int64 failed")
	ErrInterfaceToInt       = errors.New("converting interface to int failed")
)

// InterfaceToStringMap converts a interface to map[string]string.
func InterfaceToStringMap(in interface{}) (map[string]string, error) {
	if in == nil {
		return map[string]string{}, nil
	}

	m, ok := in.(map[string]interface{})

	if !ok {
		return nil, ErrInterfaceToStringMap
	}

	out := make(map[string]string)
	for key, value := range m {
		out[key] = value.(string)
	}

	return out, nil
}

// InterfaceToString converts a interface to a string.
func InterfaceToString(in interface{}) (string, error) {
	if in == nil {
		return "", nil
	}

	out, ok := in.(string)
	if !ok {
		return "", ErrInterfaceToString
	}

	return out, nil
}

// InterfaceToBool converts a interface to a bool.
func InterfaceToBool(in interface{}) (bool, error) {
	if in == nil {
		return false, nil
	}

	out, ok := in.(bool)
	if !ok {
		return false, ErrInterfaceToBool
	}

	return out, nil
}

// InterfaceToInt64 converts a interface from
// a json un-marshaled struct to a int64.
func InterfaceToInt64(in interface{}) (int64, error) {
	if in == nil {
		return 0, nil
	}

	out, ok := in.(float64)
	if !ok {
		return 0, ErrInterfaceToInt64
	}

	return int64(out), nil
}

// InterfaceToInt converts a interface from
// a json un-marshaled struct to a int.
func InterfaceToInt(in interface{}) (int, error) {
	if in == nil {
		return 0, nil
	}

	out, ok := in.(float64)
	if !ok {
		return 0, ErrInterfaceToInt
	}

	return int(out), nil
}
