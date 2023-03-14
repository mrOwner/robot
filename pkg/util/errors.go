package util

import (
	"errors"
)

func JoinErrs(err error, fn func() error) {
	err = errors.Join(err, fn())
}
