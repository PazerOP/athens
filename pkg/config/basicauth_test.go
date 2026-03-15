package config

import (
	"testing"

	"github.com/wow-look-at-my/testify/require"
)

func TestBasicAuth_Both(t *testing.T) {
	c := &Config{BasicAuthUser: "user", BasicAuthPass: "pass"}
	user, pass, ok := c.BasicAuth()
	require.True(t, ok)
	require.Equal(t, "user", user)
	require.Equal(t, "pass", pass)
}

func TestBasicAuth_NoUser(t *testing.T) {
	c := &Config{BasicAuthPass: "pass"}
	_, _, ok := c.BasicAuth()
	require.False(t, ok)
}

func TestBasicAuth_NoPass(t *testing.T) {
	c := &Config{BasicAuthUser: "user"}
	_, _, ok := c.BasicAuth()
	require.False(t, ok)
}

func TestBasicAuth_Empty(t *testing.T) {
	c := &Config{}
	_, _, ok := c.BasicAuth()
	require.False(t, ok)
}

func TestFilterOff_Empty(t *testing.T) {
	c := &Config{}
	require.True(t, c.FilterOff())
}

func TestFilterOff_WithFile(t *testing.T) {
	c := &Config{FilterFile: "/path/to/filter"}
	require.False(t, c.FilterOff())
}
