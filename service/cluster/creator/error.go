package creator

import (
	"github.com/giantswarm/microerror"
)

var invalidConfigError = microerror.New("invalid config")

// IsInvalidConfig asserts invalidConfigError.
func IsInvalidConfig(err error) bool {
	return microerror.Cause(err) == invalidConfigError
}

var invalidRequestError = microerror.New("invalid request")

// IsInvalidRequest asserts invalidRequestError.
func IsInvalidRequest(err error) bool {
	return microerror.Cause(err) == invalidRequestError
}

var notFoundError = microerror.New("not found")

// IsNotFound asserts notFoundError.
func IsNotFound(err error) bool {
	return microerror.Cause(err) == notFoundError
}
