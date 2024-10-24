package utest

import (
	errorsmod "cosmossdk.io/errors"
)

type Truer interface {
	True(value bool, msgAndArgs ...interface{})
}

func IsErr(t Truer, expected, actual error) {
	t.True(errorsmod.IsOf(actual, expected))
}
