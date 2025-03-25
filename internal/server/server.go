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

func runMiddleware(runner runnerFunc) handleFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		result := []string{}
		decorator := messagedecorator.New()
		parManager := parameters.NewHTTPPars(r)
		runner(parManager, func(eventType int, message string) {
			decoratedMessage := decorator.DecorateMessage(eventType, message)
			result = append(result, decoratedMessage)
		})

		res, err := json.Marshal(result)
		if err != nil {
			http.Error(w, fmt.Sprintf("An error occurred: %v", err), http.StatusInternalServerError)
			return
		}
		w.Write(res)
	}
}
