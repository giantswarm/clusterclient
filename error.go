package client

import (
	"fmt"

	"github.com/giantswarm/api-schema"
	"github.com/juju/errgo"
)

var (
	maskAny = errgo.MaskFunc(errgo.Any)
)

func maskAnyf(err error, f string, v ...interface{}) error {
	if err == nil {
		return nil
	}

	f = fmt.Sprintf("%s: %s", err.Error(), f)
	newErr := errgo.WithCausef(nil, errgo.Cause(err), f, v...)
	newErr.(*errgo.Err).SetLocation(1)

	return newErr
}

func mapError(err error) error {
	if apischema.IsResourceAlreadyExistsError(err) {
		return maskAny(AlreadyExistsError)
	}
	if apischema.IsResourceNotFoundError(err) {
		return maskAny(NotFoundError)
	}

	return maskAny(err)
}

var AlreadyExistsError = errgo.New("already exists")

func IsAlreadyExists(err error) bool {
	return errgo.Cause(err) == AlreadyExistsError
}

var NotFoundError = errgo.New("not found")

func IsNotFound(err error) bool {
	return errgo.Cause(err) == NotFoundError
}
