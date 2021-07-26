package task4

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFibbonaciRec(t *testing.T){
	result1 := fibRec(10)
	require.Equal(t, result1, 55)
}
