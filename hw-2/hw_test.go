package hw2

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestItoa(t *testing.T) {
	// require.Equal(t, "123", itoa(123))
	require.Equal(t, "123", tts())
}
