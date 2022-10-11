package util

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPassword(t *testing.T){
	pass := Randomstring(6)
	hashpass,err := HashPassword(pass)
	require.NoError(t,err)
	require.NotEmpty(t,hashpass)

	err = CheckPassword(pass,hashpass)
	require.NoError(t,err)
}