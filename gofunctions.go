package main

import (
	"fmt"
	"strconv"
)

func fibRec(n uint64) uint64 {
	if n <= 1 {
		return 1
	} else {
		return fibRec(n-1) + fibRec(n-2)
	}
}

func fibLoop(n int) int {
	var i, j int = 0, 1
	var c int = 1
	for k := 0; k < n; k++ {
		c = j + i
		i = j
		j = c
	}
	return c
}

func fizzbuzz(n int) string {
	if n%15 == 0 {
		return "fizzbuzz"
	} else if n%5 == 0 {
		return "buzz"
	} else if n%3 == 0 {
		return "fizz"
	} else {
		return strconv.Itoa(n)
	}
}

func palindrome(n string) {
	length := len(n)
	j := length
	status := true
	for i := 0; i < length/2; i++ {
		if n[i] != n[j-1] {
			status = false
			break
		}
		j--
	}
	if status {
		fmt.Println("Palindrome")
	} else {
		fmt.Println("Not palindrome")
	}
}

func oddevensum(n int) (int, int) {
	var sumodd, sumeven int = 0, 0
	for i := 1; i <= n; i++ {
		if i%2 == 0 {
			sumeven += i
		} else {
			sumodd += i
		}
	}
	return sumodd, sumeven
}