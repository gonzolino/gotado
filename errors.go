package gotado

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

type apiErrors struct {
	Errors []apiError `json:"errors"`
}

func (es *apiErrors) Error() string {
	errs := make([]string, len(es.Errors))
	for i, e := range es.Errors {
		errs[i] = e.Error()
	}
	return strings.Join(errs, ", ")
}

type apiError struct {
	Code  string `json:"code"`
	Title string `json:"title"`
}

func (e *apiError) Error() string {
	return fmt.Sprintf("%s: %s", strings.Title(e.Code), e.Title)
}

func isError(resp *http.Response) error {
	if resp == nil {
		return errors.New("response is nil")
	}
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		return nil
	}
	var errs apiErrors
	if err := json.NewDecoder(resp.Body).Decode(&errs); err != nil {
		return fmt.Errorf("unable to decode API error: %w", err)
	}

	if len(errs.Errors) == 1 {
		return &errs.Errors[0]
	} else if len(errs.Errors) == 0 {
		return fmt.Errorf("API returned empty error")
	} else {
		return &errs
	}
}
