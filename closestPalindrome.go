package main

import (
	"math"
	"strconv"
)

func mirror(arr []rune, chars []rune) string {
	ans := make([]rune, len(chars))
	copy(ans, chars)
	copy(ans, arr)

	mid := (len(ans) - 1) / 2
	for i := len(ans) - 1; i > mid; i-- {
		ind := len(ans) - 1 - i
		ans[i] = ans[ind]
	}
	return string(ans)
}

func NearestPalindrome(n string) string {
	chars := []rune(n)
	mid := len(chars)/2 + len(n)%2
	leftSide := chars[:mid]

	arr := []int{}
	arr = append(arr, int(math.Pow10(len(chars)-1)-1))
	for _, d := range []int{-1, 0, 1} {
		num, _ := strconv.Atoi(string(leftSide))
		num += d
		l := strconv.Itoa(num)
		c := mirror([]rune(l), chars)
		v, _ := strconv.Atoi(c)
		arr = append(arr, v)
	}

	arr = append(arr, int(math.Pow10(len(chars))+1))
	ans := 0
	diff := math.MaxInt
	orig, _ := strconv.Atoi(n)
	for _, v := range arr {
		if int(math.Abs(float64(v-orig))) < diff && v != orig {
			ans = v
			diff = int(math.Abs(float64(v - orig)))
		}
	}

	return strconv.Itoa(ans)
}
