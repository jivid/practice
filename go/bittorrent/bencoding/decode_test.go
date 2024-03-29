package bencoding

import (
	"bufio"
	"fmt"
	"math"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func newReader(t *testing.T, input string) *bufio.Reader {
	t.Helper()
	return bufio.NewReader(strings.NewReader(input))
}

func TestString(t *testing.T) {
	t.Run("Exact match", func(t *testing.T) {
		ret, err := Decode([]byte("4:name"))
		assert.Equal(t, "name", ret)
		assert.Nil(t, err)
	})

	t.Run("Stream not complete", func(t *testing.T) {
		ret, err := Decode([]byte("4:names"))
		assert.NotNil(t, err)
		assert.Nil(t, ret)
	})

	t.Run("Too few bytes", func(t *testing.T) {
		ret, err := Decode([]byte("4:nam"))
		assert.NotNil(t, err)
		assert.Nil(t, ret)
	})
}

func TestInt(t *testing.T) {
	t.Run("Signed int", func(t *testing.T) {
		ret, err := Decode([]byte("i15e"))
		assert.Equal(t, 15, ret)
		assert.Nil(t, err)
		ret, err = Decode([]byte("i-123e"))
		assert.Equal(t, -123, ret)
		assert.Nil(t, err)
	})

	t.Run("Signed int64", func(t *testing.T) {
		ret, err := Decode([]byte(fmt.Sprintf("i%de", math.MaxInt64)))
		assert.Equal(t, math.MaxInt64, ret)
		assert.Nil(t, err)
		ret, err = Decode([]byte(fmt.Sprintf("i%de", math.MinInt64)))
		assert.Equal(t, math.MinInt64, ret)
		assert.Nil(t, err)
	})

	t.Run("Stream not complete", func(t *testing.T) {
		ret, err := Decode([]byte("i3eabc"))
		assert.NotNil(t, err)
		assert.Nil(t, ret)
	})

	t.Run("Leading zero", func(t *testing.T) {
		ret, err := Decode([]byte("i03e"))
		assert.NotNil(t, err)
		assert.Nil(t, ret)
	})

	t.Run("Zero", func(t *testing.T) {
		ret, err := Decode([]byte("i0e"))
		assert.Equal(t, 0, ret)
		assert.Nil(t, err)
	})
}

func TestDict(t *testing.T) {
	t.Run("Empty dict", func(t *testing.T) {
		ret, err := Decode([]byte("de"))
		assert.Nil(t, err)
		assert.Empty(t, ret)
	})

	t.Run("Dict with nested elements", func(t *testing.T) {
		ret, err := Decode([]byte("d4:repo8:practice4:yeari2023e8:elementsl3:foo3:bare5:otherd3:bar3:bazee"))
		assert.Nil(t, err)
		if dict, ok := ret.(map[string]BencodedElement); ok {
			assert.Equal(t, "practice", dict["repo"])
			assert.Equal(t, 2023, dict["year"])
			assert.Equal(t, []BencodedElement{"foo", "bar"}, dict["elements"])
			assert.Equal(t, map[string]BencodedElement{"bar": "baz"}, dict["other"])
		} else {
			t.Fatalf("Got return that is not a map: %v", ret)
		}
	})
}
