// Package server is a HTTP server wrapper around the migrator
package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/olbrichattila/godbmigrator_cmd/internal/messagedecorator"
	"github.com/olbrichattila/godbmigrator_cmd/internal/parameters"
)

type handleFunc func(w http.ResponseWriter, r *http.Request)
type runnerFunc func(parameters parameters.ParameterManager, messageCallback func(int, string))

// Serve runs HTTP server on the specified port and call the runner which will execute migration
func Serve(port int, runner runnerFunc) error {
	http.HandleFunc("/", runMiddleware(runner))
	fmt.Printf("Listening on port %d\n", port)

	return http.ListenAndServe(":"+strconv.Itoa(port), nil)
}

type HttpResponse struct {
	Errors []string `json:"errors"`
	Body   []string `json:"body"`
}

func runMiddleware(runner runnerFunc) handleFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response := &HttpResponse{
			Errors: []string{},
			Body:   []string{},
		}
		decorator := messagedecorator.New()
		parManager := parameters.NewHTTPPars(r)
		runner(parManager, func(eventType int, message string) {
			if eventType == -2 {
				response.Errors = append(response.Errors, message)
				return
			}
			decoratedMessage := decorator.DecorateMessage(eventType, message)
			response.Body = append(response.Body, decoratedMessage)
		})

		res, err := json.Marshal(response)
		if err != nil {
			http.Error(w, fmt.Sprintf("An error occurred: %v", err), http.StatusInternalServerError)
			return
		}
		w.Write(res)
	}
}
