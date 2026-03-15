package mem

import (
	"testing"

	"github.com/wow-look-at-my/testify/require"
)

func TestNewStorage(t *testing.T) {
	s, err := NewStorage()
	require.NoError(t, err)
	require.NotNil(t, s)
}
