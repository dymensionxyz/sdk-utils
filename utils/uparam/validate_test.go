package uparam

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidateUint64(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		x := uint64(42)
		require.NoError(t, ValidateUint64(x))
	})
	t.Run("invalid", func(t *testing.T) {
		x := "foo"
		require.Error(t, ValidateUint64(x))
	})
}

func TestValidatePositiveUint64(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		x := uint64(42)
		require.NoError(t, ValidatePositiveUint64(x))
	})
	t.Run("invalid", func(t *testing.T) {
		x := uint64(0)
		require.Error(t, ValidatePositiveUint64(x))
	})
}
