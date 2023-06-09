package bencoding

import (
	"testing"
	"strings"
	"bufio"
	"math"
	"fmt"

	"github.com/stretchr/testify/assert"
)

func newReader(t *testing.T, input string) *bufio.Reader {
	t.Helper()
	return bufio.NewReader(strings.NewReader(input))
}

func TestString(t *testing.T) {
	t.Run("Exact match", func(t *testing.T) {
		ret, err := Decode(newReader(t, "4:name"))
		assert.Equal(t, "name", ret)
		assert.Nil(t, err)
	})

	t.Run("Stream not complete", func(t *testing.T) {
		ret, err := Decode(newReader(t, "4:names"))
		assert.NotNil(t, err)
		assert.Nil(t, ret)
	})

	t.Run("Too few bytes", func(t *testing.T) {
		ret, err := Decode(newReader(t, "4:nam"))
		assert.NotNil(t, err)
		assert.Nil(t, ret)
	})
}

func TestInt(t *testing.T) {
	t.Run("Signed int", func(t *testing.T) {
		ret, err := Decode(newReader(t, "i15e"))
		assert.Equal(t, 15, ret)
		assert.Nil(t, err)
		ret, err = Decode(newReader(t, "i-123e"))
		assert.Equal(t, -123, ret)
		assert.Nil(t, err)
	})

	t.Run("Signed int64", func(t *testing.T) {
		ret, err := Decode(newReader(t, fmt.Sprintf("i%de", math.MaxInt64)))
		assert.Equal(t, math.MaxInt64, ret)
		assert.Nil(t, err)
		ret, err = Decode(newReader(t, fmt.Sprintf("i%de", math.MinInt64)))
		assert.Equal(t, math.MinInt64, ret)
		assert.Nil(t, err)
	})

	t.Run("Stream not complete", func(t *testing.T) {
		ret, err := Decode(newReader(t, "i3eabc"))
		assert.NotNil(t, err)
		assert.Nil(t, ret)
	})

	t.Run("Leading zero", func(t *testing.T) {
		ret, err := Decode(newReader(t, "i03e"))
		assert.NotNil(t, err)
		assert.Nil(t, ret)
	})

	t.Run("Zero", func(t *testing.T) {
		ret, err := Decode(newReader(t, "i0e"))
		assert.Equal(t, 0, ret)
		assert.Nil(t, err)
	})
}

