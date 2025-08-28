package gotado

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/oauth2"
)

// mockHTTPClient mocks an HTTPClient by returning the stored response when Do is called.
type mockHTTPClient struct {
	Response *http.Response
	Error    error
}

// Do returns the Response stored in the mockHTTPClient.
func (c mockHTTPClient) Do(_ *http.Request) (*http.Response, error) {
	return c.Response, c.Error
}

func makeResponse(code int, body string) *http.Response {
	return &http.Response{
		StatusCode: code,
		Status:     http.StatusText(code),
		Body:       io.NopCloser(strings.NewReader(body)),
	}
}

func TestGet(t *testing.T) {
	type foobar struct {
		Foo string `json:"foo"`
		Bar string `json:"bar"`
	}

	tests := map[string]struct {
		url        string
		mockResp   *http.Response
		mockErr    error
		wantFoobar *foobar
		wantErr    error
	}{
		"Simple": {
			url:        "http://example.org",
			mockResp:   makeResponse(http.StatusOK, `{"foo": "foo","bar": "bar"}`),
			mockErr:    nil,
			wantFoobar: &foobar{Foo: "foo", Bar: "bar"},
			wantErr:    nil,
		},
		"InvalidURL": {
			url:        "invalid://url%%",
			mockResp:   nil,
			mockErr:    nil,
			wantFoobar: nil,
			wantErr:    fmt.Errorf("unable to create http request: parse \"invalid://url%%%%\": invalid URL escape \"%%%%\""),
		},
		"HTTPClientError": {
			url:        "http://example.org",
			mockResp:   nil,
			mockErr:    fmt.Errorf("http client error"),
			wantFoobar: nil,
			wantErr:    fmt.Errorf("unable to talk to tado° API: http client error"),
		},
		"UnparseableJson": {
			url:        "http://example.org",
			mockResp:   makeResponse(http.StatusOK, `{notjson}`),
			mockErr:    nil,
			wantFoobar: nil,
			wantErr:    fmt.Errorf("unable to decode tado° API response: invalid character 'n' looking for beginning of object key string"),
		},
		"EmptyErrorList": {
			url:        "http://example.org",
			mockResp:   makeResponse(http.StatusInternalServerError, `{"errors":[]}`),
			mockErr:    nil,
			wantFoobar: nil,
			wantErr:    fmt.Errorf("tado° API error: API returned empty error"),
		},
		"SingleError": {
			url:        "http://example.org",
			mockResp:   makeResponse(http.StatusInternalServerError, `{"errors":[{"code":"1","title":"One"}]}`),
			mockErr:    nil,
			wantFoobar: nil,
			wantErr:    fmt.Errorf("tado° API error: 1: One"),
		},
		"MultiError": {
			url:        "http://example.org",
			mockResp:   makeResponse(http.StatusInternalServerError, `{"errors":[{"code":"1","title":"One"},{"code":"2","title":"Two"}]}`),
			mockErr:    nil,
			wantFoobar: nil,
			wantErr:    fmt.Errorf("tado° API error: 1: One, 2: Two"),
		},
		"UnparseableError": {
			url:        "http://example.org",
			mockResp:   makeResponse(http.StatusInternalServerError, `{errorjson}`),
			mockErr:    nil,
			wantFoobar: nil,
			wantErr:    fmt.Errorf("tado° API error: unable to decode API error: invalid character 'e' looking for beginning of object key string"),
		},
	}

	ctx := context.Background()
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			client := newClient(ctx, &oauth2.Config{}, nil)
			client.http = mockHTTPClient{Response: tc.mockResp, Error: tc.mockErr}

			result := &foobar{}
			err := client.get(ctx, tc.url, result)

			if tc.wantErr != nil {
				assert.EqualError(t, err, tc.wantErr.Error())
			} else {
				assert.NoError(t, err)
				assert.Exactly(t, tc.wantFoobar, result)
			}

		})
	}
}

func TestPost(t *testing.T) {
	tests := map[string]struct {
		url      string
		mockResp *http.Response
		mockErr  error
		wantErr  error
	}{
		"Simple": {
			url:      "http://example.org",
			mockResp: makeResponse(http.StatusNoContent, ""),
			mockErr:  nil,
			wantErr:  nil,
		},
		"InvalidURL": {
			url:      "invalid://url%%",
			mockResp: nil,
			mockErr:  nil,
			wantErr:  fmt.Errorf("unable to create http request: parse \"invalid://url%%%%\": invalid URL escape \"%%%%\""),
		},
		"HTTPClientError": {
			url:      "http://example.org",
			mockResp: nil,
			mockErr:  fmt.Errorf("http client error"),
			wantErr:  fmt.Errorf("unable to talk to tado° API: http client error"),
		},
		"UnexepctedResponseCode": {
			url:      "http://example.org",
			mockResp: makeResponse(http.StatusOK, `{"foo": "foo","bar": "bar"}`),
			mockErr:  nil,
			wantErr:  fmt.Errorf("unexpected tado° API response status: OK"),
		},
		"EmptyErrorList": {
			url:      "http://example.org",
			mockResp: makeResponse(http.StatusInternalServerError, `{"errors":[]}`),
			mockErr:  nil,
			wantErr:  fmt.Errorf("tado° API error: API returned empty error"),
		},
		"SingleError": {
			url:      "http://example.org",
			mockResp: makeResponse(http.StatusInternalServerError, `{"errors":[{"code":"1","title":"One"}]}`),
			mockErr:  nil,
			wantErr:  fmt.Errorf("tado° API error: 1: One"),
		},
		"MultiError": {
			url:      "http://example.org",
			mockResp: makeResponse(http.StatusInternalServerError, `{"errors":[{"code":"1","title":"One"},{"code":"2","title":"Two"}]}`),
			mockErr:  nil,
			wantErr:  fmt.Errorf("tado° API error: 1: One, 2: Two"),
		},
		"UnparseableError": {
			url:      "http://example.org",
			mockResp: makeResponse(http.StatusInternalServerError, `{errorjson}`),
			mockErr:  nil,
			wantErr:  fmt.Errorf("tado° API error: unable to decode API error: invalid character 'e' looking for beginning of object key string"),
		},
	}

	ctx := context.Background()
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			client := newClient(ctx, &oauth2.Config{}, nil)
			client.http = mockHTTPClient{Response: tc.mockResp, Error: tc.mockErr}

			err := client.post(ctx, tc.url)

			if tc.wantErr != nil {
				assert.EqualError(t, err, tc.wantErr.Error())
			} else {
				assert.NoError(t, err)
			}

		})
	}
}

func TestPut(t *testing.T) {
	type foobar struct {
		Foo string `json:"foo"`
		Bar string `json:"bar"`
	}

	tests := map[string]struct {
		url        string
		data       interface{}
		mockResp   *http.Response
		mockErr    error
		wantFoobar interface{}
		wantErr    error
	}{
		"SimpleValue": {
			url:        "http://example.org",
			data:       foobar{Foo: "foo", Bar: "bar"},
			mockResp:   makeResponse(http.StatusOK, `{"foo": "foo","bar": "bar"}`),
			mockErr:    nil,
			wantFoobar: foobar{Foo: "foo", Bar: "bar"},
			wantErr:    nil,
		},
		"SimpleValuePtr": {
			url:        "http://example.org",
			data:       &foobar{Foo: "foo", Bar: "bar"},
			mockResp:   makeResponse(http.StatusOK, `{"foo": "foo","bar": "bar"}`),
			mockErr:    nil,
			wantFoobar: &foobar{Foo: "foo", Bar: "bar"},
			wantErr:    nil,
		},
		"InvalidURL": {
			url:        "invalid://url%%",
			data:       foobar{Foo: "foo", Bar: "bar"},
			mockResp:   nil,
			mockErr:    nil,
			wantFoobar: nil,
			wantErr:    fmt.Errorf("unable to create http request: parse \"invalid://url%%%%\": invalid URL escape \"%%%%\""),
		},
		"HTTPClientError": {
			url:        "http://example.org",
			data:       foobar{Foo: "foo", Bar: "bar"},
			mockResp:   nil,
			mockErr:    fmt.Errorf("http client error"),
			wantFoobar: nil,
			wantErr:    fmt.Errorf("unable to talk to tado° API: http client error"),
		},
		"UnparseableJson": {
			url:        "http://example.org",
			data:       foobar{Foo: "foo", Bar: "bar"},
			mockResp:   makeResponse(http.StatusOK, `{notjson}`),
			mockErr:    nil,
			wantFoobar: foobar{Foo: "foo", Bar: "bar"},
			wantErr:    nil,
		},
		"UnparseableJsonPtr": {
			url:        "http://example.org",
			data:       &foobar{Foo: "foo", Bar: "bar"},
			mockResp:   makeResponse(http.StatusOK, `{notjson}`),
			mockErr:    nil,
			wantFoobar: nil,
			wantErr:    fmt.Errorf("unable to decode tado° API response: invalid character 'n' looking for beginning of object key string"),
		},
		"EmptyErrorList": {
			url:        "http://example.org",
			data:       foobar{Foo: "foo", Bar: "bar"},
			mockResp:   makeResponse(http.StatusInternalServerError, `{"errors":[]}`),
			mockErr:    nil,
			wantFoobar: nil,
			wantErr:    fmt.Errorf("tado° API error: API returned empty error"),
		},
		"SingleError": {
			url:        "http://example.org",
			data:       foobar{Foo: "foo", Bar: "bar"},
			mockResp:   makeResponse(http.StatusInternalServerError, `{"errors":[{"code":"1","title":"One"}]}`),
			mockErr:    nil,
			wantFoobar: nil,
			wantErr:    fmt.Errorf("tado° API error: 1: One"),
		},
		"MultiError": {
			url:        "http://example.org",
			data:       foobar{Foo: "foo", Bar: "bar"},
			mockResp:   makeResponse(http.StatusInternalServerError, `{"errors":[{"code":"1","title":"One"},{"code":"2","title":"Two"}]}`),
			mockErr:    nil,
			wantFoobar: nil,
			wantErr:    fmt.Errorf("tado° API error: 1: One, 2: Two"),
		},
		"UnparseableError": {
			url:        "http://example.org",
			data:       foobar{Foo: "foo", Bar: "bar"},
			mockResp:   makeResponse(http.StatusInternalServerError, `{errorjson}`),
			mockErr:    nil,
			wantFoobar: nil,
			wantErr:    fmt.Errorf("tado° API error: unable to decode API error: invalid character 'e' looking for beginning of object key string"),
		},
		"NilData": {
			url:        "http://example.org",
			data:       nil,
			mockResp:   makeResponse(http.StatusOK, `{}`),
			mockErr:    nil,
			wantFoobar: nil,
			wantErr:    nil,
		},
		"EmptyData": {
			url:        "http://example.org",
			data:       foobar{Foo: "", Bar: ""},
			mockResp:   makeResponse(http.StatusOK, `{}`),
			mockErr:    nil,
			wantFoobar: foobar{Foo: "", Bar: ""},
			wantErr:    nil,
		},
	}

	ctx := context.Background()
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			client := newClient(ctx, &oauth2.Config{}, nil)
			client.http = mockHTTPClient{Response: tc.mockResp, Error: tc.mockErr}
			data := tc.data

			err := client.put(ctx, tc.url, data)

			if tc.wantErr != nil {
				assert.EqualError(t, err, tc.wantErr.Error())
			} else {
				assert.NoError(t, err)
				assert.Exactly(t, tc.wantFoobar, data)
			}
		})
	}
}

func TestDelete(t *testing.T) {
	tests := map[string]struct {
		url      string
		mockResp *http.Response
		mockErr  error
		wantErr  error
	}{
		"Simple": {
			url:      "http://example.org",
			mockResp: makeResponse(http.StatusNoContent, ""),
			mockErr:  nil,
			wantErr:  nil,
		},
		"UnexepctedResponseCode": {
			url:      "http://example.org",
			mockResp: makeResponse(http.StatusOK, `{"foo": "foo","bar": "bar"}`),
			mockErr:  nil,
			wantErr:  fmt.Errorf("unexpected tado° API response status: OK"),
		},
		"InvalidURL": {
			url:      "invalid://url%%",
			mockResp: nil,
			mockErr:  nil,
			wantErr:  fmt.Errorf("unable to create http request: parse \"invalid://url%%%%\": invalid URL escape \"%%%%\""),
		},
		"HTTPClientError": {
			url:      "http://example.org",
			mockResp: nil,
			mockErr:  fmt.Errorf("http client error"),
			wantErr:  fmt.Errorf("unable to talk to tado° API: http client error"),
		},
		"EmptyErrorList": {
			url:      "http://example.org",
			mockResp: makeResponse(http.StatusInternalServerError, `{"errors":[]}`),
			mockErr:  nil,
			wantErr:  fmt.Errorf("tado° API error: API returned empty error"),
		},
		"SingleError": {
			url:      "http://example.org",
			mockResp: makeResponse(http.StatusInternalServerError, `{"errors":[{"code":"1","title":"One"}]}`),
			mockErr:  nil,
			wantErr:  fmt.Errorf("tado° API error: 1: One"),
		},
		"MultiError": {
			url:      "http://example.org",
			mockResp: makeResponse(http.StatusInternalServerError, `{"errors":[{"code":"1","title":"One"},{"code":"2","title":"Two"}]}`),
			mockErr:  nil,
			wantErr:  fmt.Errorf("tado° API error: 1: One, 2: Two"),
		},
		"UnparseableError": {
			url:      "http://example.org",
			mockResp: makeResponse(http.StatusInternalServerError, `{errorjson}`),
			mockErr:  nil,
			wantErr:  fmt.Errorf("tado° API error: unable to decode API error: invalid character 'e' looking for beginning of object key string"),
		},
	}

	ctx := context.Background()
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			client := newClient(ctx, &oauth2.Config{}, nil)
			client.http = mockHTTPClient{Response: tc.mockResp, Error: tc.mockErr}

			err := client.delete(ctx, tc.url)

			if tc.wantErr != nil {
				assert.EqualError(t, err, tc.wantErr.Error())
			} else {
				assert.NoError(t, err)
			}

		})
	}
}
