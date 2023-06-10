package bencoding

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"strconv"
)

type BencodedElement interface{}

type Decoder struct {
	r   *bufio.Reader
	pos int
}

var LeadingZeroIntegerError = func(pos int) error { return fmt.Errorf("Pos: %d. Integers encoded with leading 0 are invalid", pos) }
var TooFewStringBytesError = func(pos int) error { return fmt.Errorf("Pos: %d. Didn't receive expected number of string bytes", pos) }
var StreamNotCompleteError = func(pos int) error {
	return fmt.Errorf("Pos: %d. Stream isn't complete after decoding complete element", pos)
}

func Decode(b []byte) (BencodedElement, error) {
	r := bufio.NewReader(bytes.NewReader(b))
	d := Decoder{r, 0}
	return d.Decode()
}

func (d *Decoder) Decode() (BencodedElement, error) {
	hint, err := d.r.Peek(1)
	if err != nil {
		return nil, err
	}

	switch {
	case '0' <= hint[0] && hint[0] <= '9':
		ret, err := d.decodeString()
		return d.checkStreamComplete(ret, err)
	case hint[0] == 'i':
		ret, err := d.decodeInt()
		return d.checkStreamComplete(ret, err)
	case hint[0] == 'l':
		ret, err := d.decodeList()
		return d.checkStreamComplete(ret, err)
	case hint[0] == 'd':
		ret, err := d.decodeDict()
		return d.checkStreamComplete(ret, err)
	default:
		return nil, errors.New(fmt.Sprintf("Invalid hint rune: %d", hint))
	}

}

func (d *Decoder) readByte() (byte, error) {
	b, err := d.r.ReadByte()
	d.pos += 1
	return b, err
}

func (d *Decoder) checkStreamComplete(e BencodedElement, err error) (BencodedElement, error) {
	if err != nil {
		return nil, err
	}
	if d.isStreamComplete() {
		return e, nil
	} else {
		return nil, StreamNotCompleteError(d.pos)
	}
}

func (d *Decoder) isStreamComplete() bool {
	if _, err := d.r.Peek(1); err == io.EOF {
		return true
	}

	return false
}

func (d *Decoder) decodeString() (string, error) {
	length, err := d.r.ReadString(':')
	d.pos += len(length)

	if err != nil {
		return "", err
	}

	// length will contain a trailing colon. Strip it here
	numBytes, err := strconv.Atoi(length[:len(length)-1])
	if err != nil {
		return "", err
	}
	buf := make([]byte, numBytes)
	numRead, err := d.r.Read(buf)
	d.pos += numRead
	if err != nil {
		return "", err
	}
	if numRead < numBytes {
		return "", TooFewStringBytesError(d.pos)
	}

	return string(buf), nil
}

// Assumes the leading 'i' hint has _not_ been read from the buffer yet
// Returns 0 as the integer value if an error was encountered
func (d *Decoder) decodeInt() (int, error) {
	// Drop the 'i' delimiter
	_, err := d.readByte()

	if err != nil {
		return 0, err
	}

	// Read everything until the trailing 'e' delimiter
	data, err := d.r.ReadString('e')
	d.pos += len(data)
	if err != nil {
		return 0, err
	}

	// Special case: any data with a leading 0 is invalid, except if
	// it signifies the integer 0
	if data[0] == '0' && data != "0e" {
		// We want the position in the error to be the position
		// of the 0. d.pos is current position of the reader, so
		// we need to backtrack to the zero position
		return 0, LeadingZeroIntegerError(d.pos - len(data) + 1)
	}

	// data will contain the trailing 'e'. Drop it here
	parsedInt, err := strconv.Atoi(data[:len(data)-1])
	if err != nil {
		return 0, err
	}

	return parsedInt, nil
}

func (d *Decoder) decodeList() ([]BencodedElement, error) {
	// Drop the leading 'l'
	_, err := d.readByte()
	if err != nil {
		return nil, err
	}

	buf := []BencodedElement{}

	for {
		hint, err := d.r.Peek(1)
		if err != nil {
			return nil, err
		}

		switch {
		case hint[0] == 'e':
			// Drop the trailing 'e' before yielding the stream
			_, err := d.readByte()
			if err != nil {
				return nil, err
			}
			return buf, nil
		case '0' <= hint[0] && hint[0] <= '9':
			ret, err := d.decodeString()
			if err != nil {
				return nil, err
			}
			buf = append(buf, ret)
		case hint[0] == 'i':
			ret, err := d.decodeInt()
			if err != nil {
				return nil, err
			}
			buf = append(buf, ret)
		case hint[0] == 'l':
			ret, err := d.decodeList()
			if err != nil {
				return nil, err
			}
			buf = append(buf, ret)
		case hint[0] == 'd':
			ret, err := d.decodeDict()
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

func (d *Decoder) decodeDict() (map[string]BencodedElement, error) {
	// Drop the leading 'd'
	_, err := d.readByte()
	if err != nil {
		return nil, err
	}

	buf := map[string]BencodedElement{}

	for {
		hint, err := d.r.Peek(1)
		if err != nil {
			return nil, err
		}

		if hint[0] == 'e' {
			_, err := d.readByte()
			if err != nil {
				return nil, err
			}
			return buf, nil
		}

		key, err := d.decodeString()
		if err != nil {
			return nil, err
		}

		hint, err = d.r.Peek(1)
		if err != nil {
			return nil, err
		}

		switch {
		case hint[0] == 'e':
			return nil, errors.New("Invalid end character when looking for dict value")
		case '0' <= hint[0] && hint[0] <= '9':
			value, err := d.decodeString()
			if err != nil {
				return nil, err
			}
			buf[key] = value
		case hint[0] == 'i':
			value, err := d.decodeInt()
			if err != nil {
				return nil, err
			}
			buf[key] = value
		case hint[0] == 'l':
			value, err := d.decodeList()
			if err != nil {
				return nil, err
			}
			buf[key] = value
		case hint[0] == 'd':
			value, err := d.decodeDict()
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
