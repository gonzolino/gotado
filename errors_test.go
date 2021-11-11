package gotado

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsError(t *testing.T) {
	tests := map[string]struct {
		resp    *http.Response
		wantErr error
	}{
		"NoError": {
			resp:    makeResponse(200, ""),
			wantErr: nil,
		},
		"NilError": {
			resp:    nil,
			wantErr: fmt.Errorf("response is nil"),
		},

		"NoJsonError": {
			resp:    makeResponse(404, "not found"),
			wantErr: fmt.Errorf("unable to decode API error: invalid character 'o' in literal null (expecting 'u')"),
		},
		"InvalidJsonError": {
			resp:    makeResponse(301, `{"foo": "bar"}`),
			wantErr: fmt.Errorf("API returned empty error"),
		},
		"EmptyError": {
			resp:    makeResponse(500, `{"errors":[]}`),
			wantErr: fmt.Errorf("API returned empty error"),
		},
		"SingleError": {
			resp:    makeResponse(500, `{"errors":[{"code":"1","title":"One"}]}`),
			wantErr: fmt.Errorf("1: One"),
		},
		"MultiError": {
			resp:    makeResponse(500, `{"errors":[{"code":"1","title":"One"},{"code":"2","title":"Two"}]}`),
			wantErr: fmt.Errorf("1: One, 2: Two"),
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			err := isError(tc.resp)

			if tc.wantErr != nil {
				assert.EqualError(t, err, tc.wantErr.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
