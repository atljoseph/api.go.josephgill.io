package requester

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/atljoseph/api.josephgill.io/apierr"
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
// TODO: Extract slice-valued query params. use this func, or another func
func GetQueryParams(r *http.Request, qps ...*QueryParam) error {
	funcTag := "GetQueryParams"

	pkgLog.WithFunc(funcTag).WithMessage("getting query parameter values")

	// get the query params passed in with the request
	// this can include query params, if required in the route config, BUT NOT optional ones
	requestQueryParams := r.URL.Query()

	// loop through all the query params we care about
	for _, qp := range qps {

		val := requestQueryParams[qp.Name]

		// if required and no value, return error
		if qp.Required && len(val) == 0 {
			err := fmt.Errorf("no value for required query parameter")
			return apierr.Errorf(err, funcTag, "key: %s", qp.Name)
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
