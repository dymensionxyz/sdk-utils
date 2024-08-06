package uparam

import (
	errorsmod "cosmossdk.io/errors"
	"github.com/dymensionxyz/gerr-cosmos/gerrc"
)

func Uint64(x any) (uint64, error) {
	v, ok := x.(uint64)
	if !ok {
		return 0, errorsmod.WithType(gerrc.ErrInvalidArgument, x)
	}
	return v, nil
}

func ValidateUint64(x any) error {
	_, err := Uint64(x)
	return err
}

func ValidatePositiveUint64(x any) error {
	y, err := Uint64(x)
	if err != nil {
		return err
	}
	if y == 0 {
		return gerrc.ErrOutOfRange.Wrap("must be positive")
	}
	return nil
}
