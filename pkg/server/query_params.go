package server

import (
	"net/http"

	"github.com/datasweet/cast"
)

// Query gets a query param
func QueryParam(r *http.Request, key string) (string, bool) {
	if values, ok := r.URL.Query()[key]; ok && len(values) == 1 {
		return values[0], true
	}
	return "", false
}

// QueryString gets a string query param
func QueryString(r *http.Request, key string, defaultValue ...string) string {
	if p, ok := QueryParam(r, key); ok {
		return p
	}
	return defaultString(defaultValue...)
}

// QueryInt gets an int query param
func QueryInt(r *http.Request, key string, defaultValue ...int) int {
	if p, ok := QueryParam(r, key); ok {
		if n, ok := cast.AsInt(p); ok {
			return n
		}
	}
	return defaultInt(defaultValue...)
}

// TODO int8, int16, int64, float32, float64 etc ?

// defaulting
func defaultString(val ...string) string {
	if len(val) > 0 {
		return val[0]
	}
	return ""
}

func defaultInt(val ...int) int {
	if len(val) > 0 {
		return val[0]
	}
	return 0
}
