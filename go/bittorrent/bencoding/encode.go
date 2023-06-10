package bencoding

import (
	"strconv"
	"reflect"
	"fmt"
	"sort"
)

var CannotEncodeType = func(v reflect.Value, k reflect.Kind) error {return fmt.Errorf("Cannot encode value %v of kind %s", v, k)}
type encoderFunc func(reflect.Value) ([]byte, error)

func Encode(input any) ([]byte, error) {
	encoder := getEncoder(reflect.TypeOf(input))
	return encoder(reflect.ValueOf(input))
}

func getEncoder(t reflect.Type) encoderFunc {
	switch t.Kind() {
	case reflect.String:
		return encodeString
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return encodeInt
	case reflect.Interface:
		return encodeInterface
	case reflect.Array, reflect.Slice:
		return encodeList
	case reflect.Map:
		return encodeMap
	default:
		return encodeUnsupported
	}
}

func encodeString(v reflect.Value) ([]byte, error) {
	s := v.String()
	return []byte(strconv.Itoa(len(s))+":"+s), nil
}

func encodeInt(v reflect.Value) ([]byte, error) {
	return []byte("i"+strconv.FormatInt(v.Int(), 10)+"e"), nil
}

func encodeInterface(v reflect.Value) ([]byte, error) {
	return getEncoder(v.Elem().Type())(v.Elem())
}

func encodeList(v reflect.Value) ([]byte, error) {
	ret := []byte("l")
	for i := 0; i < v.Len(); i++ {
		encoded, err := getEncoder(v.Index(i).Type())(v.Index(i))
		if err != nil {
			return nil, err
		}
		ret = append(ret, encoded...)
	}
	ret = append(ret, []byte("e")...)
	return ret, nil
}

func sortStringValues(values []reflect.Value) []reflect.Value {
	strings := []string{}
	for _, v := range values {
		strings = append(strings, v.String())
	}
	sort.Strings(strings)
	vals := []reflect.Value{}
	for _, s := range strings {
		vals = append(vals, reflect.ValueOf(s))
	}
	return vals
}

func encodeMap(v reflect.Value) ([]byte, error) {
	ret := []byte("d")
	// This will panic if key is not a string, but our protocol
	// requires that map keys be strings, so allow this panic
	// TODO: Recover the panic and return a useful error
	keys := sortStringValues(v.MapKeys())
	for _, k := range keys {
		encoded, err := getEncoder(k.Type())(k)
		if err != nil {
			return nil, err
		}
		ret = append(ret, encoded...)
		fmt.Println(string(ret))

		val := v.MapIndex(k)
		encoded, err = getEncoder(val.Type())(val)
		if err != nil {
			return nil, err
		}
		ret = append(ret, encoded...)
		fmt.Println(string(ret))
	}
	ret = append(ret, []byte("e")...)
	fmt.Println(string(ret))
	return ret, nil
}

func encodeUnsupported(v reflect.Value) ([]byte, error) {
	return nil, CannotEncodeType(v, v.Kind())
}
