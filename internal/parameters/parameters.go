// Package parameters handling parameters from command line or HTTP
package parameters

// ParameterManager to encapsulate parameter handling
type ParameterManager interface {
	Command() string
	Params() []string
}
