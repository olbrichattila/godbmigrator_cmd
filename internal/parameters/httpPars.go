package parameters

import (
	"net/http"
	"strconv"
	"strings"
)

// ForceParam is the parameter to force migration regardless of validation error
const ForceParam = "-force"

// NewHTTPPars create a parameter manager which works with HTTP requests
func NewHTTPPars(r *http.Request) ParameterManager {
	return &httpPars{
		r: r,
	}
}

type httpPars struct {
	r *http.Request
}

// Command implements ParameterManager.
func (h *httpPars) Command() string {
	return h.r.URL.Query().Get("command")
}

// Params implements ParameterManager, convert URL params like command line params
func (h *httpPars) Params() []string {
	if h.Command() == "add" {
		name := h.r.URL.Query().Get("name")
		if name != "" {
			return []string{name}
		}

		return []string{}
	}

	result := []string{
		h.getCountWithDefault(),
	}

	if h.getForce() {
		result = append(result, ForceParam)
	}

	return result
}

func (h *httpPars) getCountWithDefault() string {
	defaultCount := "0"
	countStr := h.r.URL.Query().Get("count")
	if countStr == "" {
		return defaultCount
	}

	if _, err := strconv.Atoi(countStr); err == nil {
		return countStr
	}

	return defaultCount
}

func (h *httpPars) getForce() bool {
	forceStr := strings.ToLower(h.r.URL.Query().Get("force"))
	return forceStr == "1" || forceStr == "true"
}
