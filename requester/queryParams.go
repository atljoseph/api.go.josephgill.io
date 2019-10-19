package requester

import (
	"fmt"
	"net/http"
	"strings"
)

// QueryParam represents a query parameter and its default value
type QueryParam struct {
	Name         string
	DefaultValue string
	Required     bool
	Value        string
	// ValueAsSlice []string
	// IsSlice defaults to false. If true, value will be passed as ValueAsSlice instead of
	// IsSlice bool
}

// GetQueryParams extracts query parameter values from the request
// needs work in order to extract slice-valued params
func GetQueryParams(r *http.Request, qps ...*QueryParam) error {
	errTag := "requester.GetQueryParams"

	// get the query params passed in with the request
	requestQueryParams := r.URL.Query()

	// loop through all the query params we care about
	for _, qp := range qps {

		val := requestQueryParams[qp.Name]

		// if required and no value, return error
		if qp.Required && len(val) == 0 {
			return fmt.Errorf("%s: no value for required key (%s)", errTag, qp.Name)
		}

		// if no value, set a default
		if len(val) == 0 {
			val = []string{qp.DefaultValue}
		}

		// need to change this
		singleVal := val[0]

		// sanitized the value
		newValWithSlashes := strings.Replace(singleVal, "_slash_", "/", -1)
		qp.Value = newValWithSlashes
	}

	return nil
}
