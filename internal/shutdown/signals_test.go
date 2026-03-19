//go:build unix

package shutdown

import (
	"testing"

	"github.com/wow-look-at-my/testify/require"
)

func TestGetSignals(t *testing.T) {
	signals := GetSignals()
	require.Len(t, signals, 2)
}
