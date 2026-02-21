package server

import (
	"encoding/json"
	"net/http"
	"regexp"
	"strconv"
	"time"

	"mirage/src/models"
)

// pathVarRe matches {varName} in a path and captures the name (e.g. "{id}" → "id").
var pathVarRe = regexp.MustCompile(`\{([^}]+)\}`)

// Returns the URL path parameter values as a map.
func pathParams(r *http.Request, path string) map[string]any {
	matches := pathVarRe.FindAllStringSubmatch(path, -1)
	if len(matches) == 0 {
		return nil
	}
	params := make(map[string]any)
	for _, m := range matches {
		if len(m) > 1 {
			params[m[1]] = asJSONType(r.PathValue(m[1]))
		}
	}
	return params
}

// Converts the raw path value to the usual REST/JSON type:
// integer and decimal -> number, otherwise -> string.
func asJSONType(s string) any {
	if n, err := strconv.ParseInt(s, 10, 64); err == nil {
		return n
	}
	if f, err := strconv.ParseFloat(s, 64); err == nil {
		return f
	}
	return s
}

// Replaces placeholders in the response with the URL path parameter values.
func injectPathParams(v any, params map[string]any) any {
	if len(params) == 0 {
		return v
	}
	switch x := v.(type) {
	case string:
		if len(x) > 2 && x[0] == '{' && x[len(x)-1] == '}' {
			if val, ok := params[x[1:len(x)-1]]; ok {
				return val
			}
		}
		return x
	case map[string]any:
		out := make(map[string]any, len(x))
		for k, val := range x {
			out[k] = injectPathParams(val, params)
		}
		return out
	case []any:
		out := make([]any, len(x))
		for i, val := range x {
			out[i] = injectPathParams(val, params)
		}
		return out
	}
	return v
}

func WriteResponse(w http.ResponseWriter, ep *models.Endpoint, r *http.Request) {
	if ep.Delay != nil {
		time.Sleep(time.Duration(*ep.Delay) * time.Millisecond)
	}

	res := ep.Response
	if r != nil {
		// If the path has variables (e.g. /users/{id}), replace "{id}" in the response with request values.
		if params := pathParams(r, ep.Path); len(params) > 0 {
			res = injectPathParams(ep.Response, params)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	if ep.Status != nil {
		w.WriteHeader(*ep.Status)
	}
	json.NewEncoder(w).Encode(res)
}
