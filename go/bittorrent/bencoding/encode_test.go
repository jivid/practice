package bencoding

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestEncodeString(t *testing.T) {
	ret, _ := Encode("input")
	assert.Equal(t, []byte("5:input"), ret)
}

func TestEncodeInt(t *testing.T) {
	ret, _ := Encode(15)
	assert.Equal(t, []byte("i15e"), ret)
}

func TestEncodeList(t *testing.T) {
	ret, err := Encode([]interface{}{"foo", "bar", 1, map[string]interface{}{"bar": []interface{}{1,2,3}}})
	assert.Equal(t, []byte("l3:foo3:bari1ed3:barli1ei2ei3eeee"), ret)
	assert.Nil(t, err)
}

func TestEncodeDict(t *testing.T) {
	ret, err := Encode(map[string]interface{}{"repo": "practice", "year": 2023, "languages": []string{"golang"}})
	assert.Equal(t, []byte("d9:languagesl6:golange4:repo8:practice4:yeari2023ee"), ret)
	assert.Nil(t, err)
}
