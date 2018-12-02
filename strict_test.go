package strict

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

const UnsetVar = "BABABA_BA_BANANA"

func TestMustGetenv(t *testing.T) {
	assert.Panics(t, func() { MustGetenv(UnsetVar) })
}

func TestGetenv(t *testing.T) {
	value, err := Getenv(UnsetVar)
	require.Equal(t, value, "")
	require.NotNil(t, err)
}

func TestGetenvWithDefault(t *testing.T) {
	value := GetenvWithDefault(UnsetVar, "le minion")
	require.Equal(t, value, "le minion")
}
