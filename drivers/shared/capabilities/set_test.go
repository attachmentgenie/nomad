package capabilities

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSet_Empty(t *testing.T) {
	t.Parallel()

	t.Run("nil", func(t *testing.T) {
		result := New(nil).Empty()
		require.True(t, result)
	})

	t.Run("empty", func(t *testing.T) {
		result := New([]string{}).Empty()
		require.True(t, result)
	})

	t.Run("full", func(t *testing.T) {
		result := New([]string{"chown", "sys_time"}).Empty()
		require.False(t, result)
	})
}
