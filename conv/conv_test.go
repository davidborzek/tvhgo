package conv_test

import (
	"testing"

	"github.com/davidborzek/tvhgo/conv"
	"github.com/stretchr/testify/assert"
)

func TestInterfaceToStringMap(t *testing.T) {
	out, err := conv.InterfaceToStringMap(
		map[string]interface{}{
			"key": "value",
		},
	)

	assert.Nil(t, err)
	assert.Equal(t, map[string]string{
		"key": "value",
	}, out)
}

func TestInterfaceToStringMapReturnsError(t *testing.T) {
	out, err := conv.InterfaceToStringMap("invalid")

	assert.Nil(t, out)
	assert.Equal(t, conv.ErrInterfaceToStringMap, err)
}

func TestInterfaceToStringMapReturnsEmptyMapForNilInput(t *testing.T) {
	out, err := conv.InterfaceToStringMap(nil)

	assert.Nil(t, err)
	assert.Empty(t, out)
}

func TestInterfaceToString(t *testing.T) {
	out, err := conv.InterfaceToString("someString")

	assert.Nil(t, err)
	assert.Equal(t, "someString", out)
}

func TestInterfaceToStringReturnsError(t *testing.T) {
	out, err := conv.InterfaceToString(1234)

	assert.Empty(t, out)
	assert.Equal(t, conv.ErrInterfaceToString, err)
}

func TestInterfaceToStringReturnsEmptyStringForNilInput(t *testing.T) {
	out, err := conv.InterfaceToString(nil)

	assert.Nil(t, err)
	assert.Equal(t, "", out)
}

func TestInterfaceToBool(t *testing.T) {
	out, err := conv.InterfaceToBool(true)

	assert.Nil(t, err)
	assert.Equal(t, true, out)
}

func TestInterfaceToBoolReturnsError(t *testing.T) {
	out, err := conv.InterfaceToBool("invalid")

	assert.False(t, out)
	assert.Equal(t, conv.ErrInterfaceToBool, err)
}

func TestInterfaceToBoolReturnsFalseForNilInput(t *testing.T) {
	out, err := conv.InterfaceToBool(nil)

	assert.Nil(t, err)
	assert.Equal(t, false, out)
}

func TestInterfaceToInt64(t *testing.T) {
	out, err := conv.InterfaceToInt64(float64(1234))

	assert.Nil(t, err)
	assert.Equal(t, int64(1234), out)
}

func TestInterfaceToInt64ReturnsError(t *testing.T) {
	out, err := conv.InterfaceToInt64("invalid")

	assert.Equal(t, int64(0), out)
	assert.Equal(t, conv.ErrInterfaceToInt64, err)
}

func TestInterfaceToInt64Returns0ForNilInput(t *testing.T) {
	out, err := conv.InterfaceToInt64(nil)

	assert.Nil(t, err)
	assert.Equal(t, int64(0), out)
}

func TestInterfaceToInt(t *testing.T) {
	out, err := conv.InterfaceToInt(float64(1234))

	assert.Nil(t, err)
	assert.Equal(t, 1234, out)
}

func TestInterfaceToIntReturnsError(t *testing.T) {
	out, err := conv.InterfaceToInt("invalid")

	assert.Equal(t, 0, out)
	assert.Equal(t, conv.ErrInterfaceToInt, err)
}

func TestInterfaceToIntReturns0ForNilInput(t *testing.T) {
	out, err := conv.InterfaceToInt(nil)

	assert.Nil(t, err)
	assert.Equal(t, 0, out)
}
