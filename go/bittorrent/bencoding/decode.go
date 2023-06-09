package bencoding

import (
	"bufio"
	"io"
	// "bytes"
	"errors"
	"fmt"
	"strconv"
)

type BencodedElement interface{}

// TODO: Enhance with positions in the stream
var LeadingZeroIntegerError = errors.New("Integers encoded with leading 0 are invalid")
var TooFewStringBytesError = errors.New("Didn't receive expected number of string bytes")
var StreamNotCompleteError = errors.New("Stream isn't complete after decoding complete element")

func Decode(r *bufio.Reader) (BencodedElement, error) {
	hint, err := r.Peek(1)
	if err != nil {
		return nil, err
	}

	switch {
	case '0' <= hint[0] && hint[0] <= '9':
		ret, err := DecodeString(r)
		return CheckStreamComplete(r, ret, err)
	case hint[0] == 'i':
		ret, err := DecodeInt(r)
		return CheckStreamComplete(r, ret, err)
	case hint[0] == 'l':
		ret, err := DecodeList(r)
		return CheckStreamComplete(r, ret, err)
	case hint[0] == 'd':
		ret, err := DecodeDict(r)
		return CheckStreamComplete(r, ret, err)
	default:
		return nil, errors.New(fmt.Sprintf("Invalid hint rune: %d", hint))
	}
}

func CheckStreamComplete(r *bufio.Reader, e BencodedElement, err error) (BencodedElement, error) {
	if err != nil {
		return nil, err
	}
	if IsStreamComplete(r) {
		return e, nil
	} else {
		return nil, StreamNotCompleteError
	}
}

func IsStreamComplete(r *bufio.Reader) bool {
	if _, err := r.Peek(1); err == io.EOF {
		return true
	}

	return false
}

func DecodeString(r *bufio.Reader) (string, error) {
	length, err := r.ReadString(':')

	if err != nil {
		return "", err
	}

	// length will contain a trailing colon. Strip it here
	numBytes, err := strconv.Atoi(length[:len(length)-1])
	if err != nil {
		return "", err
	}
	buf := make([]byte, numBytes)
	numRead, err := r.Read(buf)
	if err != nil {
		return "", err
	}
	if numRead < numBytes {
		return "", TooFewStringBytesError
	}

	return string(buf), nil
}

// Assumes the leading 'i' hint has _not_ been read from the buffer yet
// Returns 0 as the integer value if an error was encountered
func DecodeInt(r *bufio.Reader) (int, error) {
	// Drop the 'i' delimiter
	_, err := r.ReadByte()
	if err != nil {
		return 0, err
	}

	// Read everything until the trailing 'e' delimiter
	data, err := r.ReadString('e')
	if err != nil {
		return 0, err
	}

	// Special case: any data with a leading 0 is invalid, except if
	// it signifies the integer 0
	if data[0] == '0' && data != "0e" {
		return 0, LeadingZeroIntegerError
	}

	// data will contain a trailing colon. Strip it here
	parsedInt, err := strconv.Atoi(data[:len(data)-1])
	if err != nil {
		return 0, err
	}

	return parsedInt, nil
}

func DecodeList(r *bufio.Reader) ([]BencodedElement, error) {
	// Drop the leading 'l'
	_, err := r.ReadByte()
	if err != nil {
		return nil, err
	}

	buf := []BencodedElement{}

	for {
		hint, err := r.Peek(1)
		if err != nil {
			return nil, err
		}

		switch {
		case hint[0] == 'e':
			// Drop the trailing 'e' before yielding the stream
			_, err := r.ReadByte()
			if err != nil {
				return nil, err
			}
			return buf, nil
		case '0' <= hint[0] && hint[0] <= '9':
			ret, err := DecodeString(r)
			if err != nil {
				return nil, err
			}
			buf = append(buf, ret)
		case hint[0] == 'i':
			ret, err := DecodeInt(r)
			if err != nil {
				return nil, err
			}
			buf = append(buf, ret)
		case hint[0] == 'l':
			ret, err := DecodeList(r)
			if err != nil {
				return nil, err
			}
			buf = append(buf, ret)
		case hint[0] == 'd':
			ret, err := DecodeDict(r)
			if err != nil {
				return nil, err
			}
			buf = append(buf, ret)
		default:
			return nil, errors.New(fmt.Sprintf("Invalid hint rune: %d", hint))
		}
	}
	return buf, nil
}

func DecodeDict(r *bufio.Reader) (map[string]BencodedElement, error) {
	// Drop the leading 'd'
	_, err := r.ReadByte()
	if err != nil {
		return nil, err
	}

	buf := map[string]BencodedElement{}

	for {
		hint, err := r.Peek(1)
		if err != nil {
			return nil, err
		}

		if hint[0] == 'e' {
			_, err := r.ReadByte()
			if err != nil {
				return nil, err
			}
			return buf, nil
		}

		key, err := DecodeString(r)
		if err != nil {
			return nil, err
		}

		hint, err = r.Peek(1)
		if err != nil {
			return nil, err
		}

		switch {
		case hint[0] == 'e':
			return nil, errors.New("Invalid end character when looking for dict value")
		case '0' <= hint[0] && hint[0] <= '9':
			value, err := DecodeString(r)
			if err != nil {
				return nil, err
			}
			buf[key] = value
		case hint[0] == 'i':
			value, err := DecodeInt(r)
			if err != nil {
				return nil, err
			}
			buf[key] = value
		case hint[0] == 'l':
			value, err := DecodeList(r)
			if err != nil {
				return nil, err
			}
			buf[key] = value
		case hint[0] == 'd':
			value, err := DecodeDict(r)
			if err != nil {
				return nil, err
			}
			buf[key] = value
		default:
			return nil, errors.New(fmt.Sprintf("Invalid hint rune: %d", hint))
		}

	}

	return buf, nil
}
