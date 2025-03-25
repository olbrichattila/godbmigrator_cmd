package parameters

import (
	"os"
)

// NewCommandLine creates an instance of parameter manager working with command line parameters
func NewCommandLine() ParameterManager {
	return &cLine{}
}

type cLine struct {
}

// Command implements ParameterManager.
func (c *cLine) Command() string {
	if len(os.Args) < 2 {
		return ""
	}
	return os.Args[1]
}

// Param implements ParameterManager.
func (c *cLine) Params() []string {
	if len(os.Args) < 2 {
		return []string{}
	}

	return os.Args[2:]
}
