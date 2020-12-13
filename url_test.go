package gotado

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestApiURL(t *testing.T) {
	tests := []struct {
		format   string
		a        []interface{}
		expected string
	}{
		{"", []interface{}{}, "https://my.tado.com/api/v2"},
		{"/", []interface{}{}, "https://my.tado.com/api/v2"},
		{"test", []interface{}{}, "https://my.tado.com/api/v2/test"},
		{"/test", []interface{}{}, "https://my.tado.com/api/v2/test"},
		{"/a/%s/", []interface{}{"b"}, "https://my.tado.com/api/v2/a/b"},
		{"a/%d", []interface{}{1}, "https://my.tado.com/api/v2/a/1"},
		{"a/%s/b/%d/c", []interface{}{"x", 1}, "https://my.tado.com/api/v2/a/x/b/1/c"},
	}

	for _, test := range tests {
		actual := apiURL(test.format, test.a...)
		assert.Exactly(t, test.expected, actual)
	}
}
