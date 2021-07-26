package task4

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFibRec(t *testing.T){
	result1 := fibRec(10)
	require.Equal(t, result1, 89)
	result2 := fibRec(11)
	require.NotEqual(t, result2, 89)
}

func TestFibLoop(t *testing.T){
	result1 := fibLoop(10)
	require.Equal(t, result1, 89)
	result2 := fibLoop(11)
	require.NotEqual(t, result2, 140)
}

func TestFizzBuzz(t *testing.T){
	result1 := fizzbuzz(45)
	require.Equal(t, result1, "fizzbuzz")
	require.NotEqual(t, result1, "fizz")
	require.NotEqual(t, result1, "buzz")
	require.NotEqual(t, result1, "15")

	result2 := fizzbuzz(55)
	require.NotEqual(t, result2, "fizzbuzz")
	require.NotEqual(t, result2, "fizz")
	require.Equal(t, result2, "buzz")
	require.NotEqual(t, result2, "45")

	result2 = fizzbuzz(81)
	require.NotEqual(t, result2, "fizzbuzz")
	require.Equal(t, result2, "fizz")
	require.NotEqual(t, result2, "buzz")
	require.NotEqual(t, result2, "81")

	result2 = fizzbuzz(64)
	require.NotEqual(t, result2, "fizzbuzz")
	require.NotEqual(t, result2, "fizz")
	require.NotEqual(t, result2, "buzz")
	require.Equal(t, result2, "64")
	require.NotEqual(t, result2, "80")
}

func TestPalindrome(t *testing.T){
	result := palindrome("sanjar")
	require.False(t, result, true)

	result = palindrome("pallap")
	require.True(t, result)
}