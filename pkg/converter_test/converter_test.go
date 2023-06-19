package converter_test

import (
	"testing"

	"github.com/pericles-luz/go-converter/pkg/converter"
	"github.com/stretchr/testify/require"
)

func TestMapInterfaceToByteMustRunCorretly(t *testing.T) {
	// Given
	data := map[string]interface{}{
		"cpf":       1,
		"sentMedia": 2,
	}
	// When
	bytes := converter.MapInterfaceToBytes(data)
	// Then
	require.NotEmpty(t, bytes)
}

func TestMapInterfaceToByteMustReturnNilIfDataIsNil(t *testing.T) {
	// Given
	var data map[string]interface{}
	// When
	bytes := converter.MapInterfaceToBytes(data)
	// Then
	require.Nil(t, bytes)
}

func TestByteToMapInterfaceMustRunCorretly(t *testing.T) {
	// Given
	data := map[string]interface{}{
		"cpf":       1,
		"sentMedia": 2,
	}
	bytes := converter.MapInterfaceToBytes(data)
	// When
	data2 := converter.ByteToMapInterface(bytes)
	// Then
	require.NotEmpty(t, data2)
}

func TestByteToMapInterfaceMustReturnNilIfBytesIsNil(t *testing.T) {
	// Given
	var bytes []byte
	// When
	data := converter.ByteToMapInterface(bytes)
	// Then
	require.Nil(t, data)
}

func TestByteToMapInterfaceMustReturnNilIfBytesIsEmpty(t *testing.T) {
	// Given
	bytes := []byte{}
	// When
	data := converter.ByteToMapInterface(bytes)
	// Then
	require.Nil(t, data)
}

func TestByteToMapInterfaceMustReturnNilIfBytesIsNotJson(t *testing.T) {
	// Given
	bytes := []byte("not json")
	// When
	data := converter.ByteToMapInterface(bytes)
	// Then
	require.Nil(t, data)
}

func TestStringToIntMustConvertCorrectly(t *testing.T) {
	// Given
	value := "123"
	// When
	result := converter.StringToInt(value)
	// Then
	require.Equal(t, 123, result)
}

func TestStringToIntMustReturnZeroIfValueIsNotNumber(t *testing.T) {
	// Given
	value := "not number"
	// When
	result := converter.StringToInt(value)
	// Then
	require.Equal(t, 0, result)
}

func TestInterfaceToIntMustConvertStringToInt(t *testing.T) {
	// Given
	value := "123"
	// When
	result := converter.InterfaceToInt(value)
	// Then
	require.Equal(t, 123, result)
}

func TestInterfaceToIntMustConvertIntToInt(t *testing.T) {
	// Given
	value := 123
	// When
	result := converter.InterfaceToInt(value)
	// Then
	require.Equal(t, 123, result)
}

func TestWhatsappNumberToBrazilianPhonenumber(t *testing.T) {
	// Given
	number := "553186058910"
	// When
	result := converter.WhatsappNumberToBrazilianPhonenumber(number)
	// Then
	require.Equal(t, "31986058910", result)
}

func TestIntToString(t *testing.T) {
	// Given
	value := 123
	// When
	result := converter.IntToString(value)
	// Then
	require.Equal(t, "123", result)
}
