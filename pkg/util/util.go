package util

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// A simple, generic function that returns the sum of an integer slice.
func Sumint(slice []int) int {
	result := 0

	for _, n := range slice {
		result += n
	}

	return result
}

// A simple generic function that returns the sum of a float64 slice.
func Sumfl(slice []float64) float64 {
	result := 0.0
	for _, n := range slice {
		result += n
	}
	return math.Floor(result*100) / 100
}

// A simple generic function that returns the sum of a float32 slice.
func Sumfl32(slice []float32) float32 {
	result := float32(0.0)
	for _, n := range slice {
		result += n
	}

	return result
}

// Returns a string value. If the number is negative, it will be wrapped in parenthesis
// i.e. -1 would be (1)
func Negint(n int) string {
	if n < 0 {
		formatted := "(" + strings.Replace(strconv.Itoa(n), "-", "", 1) + ")"
		return formatted
	} else {
		return strconv.Itoa(n)
	}
}

// Returns a string value. If the float is negative, it will be wrapped in parenthesis
// i.e. -2.5 would be (2.5)
func Negfl(n float64) string {
	if n < 0 {
		str := fmt.Sprintf("%f", math.Floor(n*100)/100)
		formatted := "(" + strings.Replace(str, "-", "", 1) + ")"
		return formatted
	} else {
		return fmt.Sprintf("%f", math.Floor(n*100)/100)
	}
}

// Returns the average of an integer slice.
func Average(slice []int) float64 {
	result := 0.0

	for _, n := range slice {
		result += float64(n)
	}

	return math.Round(result/float64(len(slice))*100) / 100
}

// Returns the middle most number in an integer slice.
func Median(slice []int) []int {
	md := []int{}

	if len(slice)%2 == 0 {
		md = append(md, slice[len(slice)/2-1])
		md = append(md, slice[len(slice)/2])
	} else {
		md = append(md, slice[len(slice)/2])
	}

	return md
}

// Returns the most repeated number in an integer slice. It will return multiple numbers if there is a tie.
func Mode(slice []int) []int {
	m := map[int]int{}
	md := []int{}
	var max int

	for i := 0; i < len(slice); i++ {
		if m[slice[i]] == m[slice[i]] {
			m[slice[i]]++
			if m[slice[i]] > max {
				max = m[slice[i]]
			}

		}
	}
	for v := range m {
		if m[v] >= max {
			md = append(md, v)
		}
	}
	return md
}

// Returns the most repeated number in an float64 slice. It will return multiple numbers if there is a tie.
func Modefl64(slice []float64) []float64 {
	m := map[float64]int{}
	md := []float64{}
	var max int

	for i := 0; i < len(slice); i++ {
		if m[slice[i]] == m[slice[i]] {
			m[slice[i]]++
			if m[slice[i]] > max {
				max = m[slice[i]]
			}
		}
	}
	for v := range m {
		if m[v] >= max {
			md = append(md, v)
		}
	}
	return md
}

// Returns the most repeated string in an string slice. It will return multiple strings if there is a tie.

func Modestr(slice []string) []string {
	m := map[string]int{}
	md := []string{}
	var max int

	for i := 0; i < len(slice); i++ {
		if m[slice[i]] == m[slice[i]] {
			m[slice[i]]++
			if m[slice[i]] > max {
				max = m[slice[i]]
			}
		}
	}
	for v := range m {
		if m[v] >= max {
			md = append(md, v)
		}
	}
	return md
}

// Returns a boolean value if the number is even.
func IsEven(n int) bool {
	if n%2 == 0 {
		return true
	} else {
		return false
	}
}

// Returns a slice of all prime numbers in a given range.
// i.e. between 2 and 10, [2, 3, 5, 7] would be the returned slice.
func Prime(start, end int) []int {
	var primes []int
	if start < 2 || end < 2 {
		return primes
	}

	for start <= end {
		isPrime := true
		for i := 2; i <= int(math.Sqrt(float64(start))); i++ {
			if start%i == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			primes = append(primes, start)
		}
		start++
	}
	return primes
}

// Returns an integer slice sorted from least to greatest.
func Ascend(arr []int) []float64 {
	slice := make([]float64, len(arr))

	for i, v := range arr {
		slice[i] = float64(v)
	}
	for i := 0; i < len(slice)-1; i++ {
		for j := 0; j < len(slice)-i-1; j++ {
			if slice[j] > slice[j+1] {
				slice[j], slice[j+1] = slice[j+1], slice[j]
			}
		}
	}

	return slice
}

// Returns an integer slice sorted from greatest to least.
func Descend(arr []int) []float64 {
	slice := make([]float64, len(arr))
	for i, v := range arr {
		slice[i] = float64(v)
	}
	for i := 0; i < len(slice)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if slice[j] < slice[j+1] {
				slice[j], slice[j+1] = slice[j+1], slice[j]
			}
		}
	}

	return slice
}

// Returns the slice index of a given string.
func IndexOf(v string, slice []string) int {
	for i := 0; i < len(slice); i++ {
		if slice[i] == v {
			return i
		}
	}

	return -1
}

// Returns the slice index of a given integer.
func IndexOfint(v int, slice []int) int {
	for i := 0; i < len(slice); i++ {
		if slice[i] == v {
			return i
		}
	}

	return -1
}

// Returns the slice index of a given float64 value.
func IndexOffl64(v float64, slice []float64) int {
	for i := 0; i < len(slice); i++ {
		if slice[i] == v {
			return i
		}
	}
	return -1
}

// Reverses a single string, with no spaces.
// i.e. apple would be elppa, the brown fox jumps over the lazy dog would be godyzalehtrevospmujxofnworbeht
func Reverse(str string) string {
	slice := strings.Split(str, "")
	r := []string{}

	for i := len(slice) - 1; i >= 0; i-- {
		r = append(r, slice[i])
	}

	return strings.Join(r, "")
}

// Reverse a single string, but handles spaces.
func ReverseLrg(str string) string {
	slice := strings.Split(str, "")
	r := []string{}

	for i := len(slice) - 1; i >= 0; i-- {
		w := strings.Split(slice[i], "")
		rw := []string{}

		for j := len(w) - 1; j >= 0; j-- {
			rw = append(rw, w[j])
		}
		r = append(r, strings.Join(rw, ""))
	}

	return strings.Join(r, "")
}

// Reverses the string, but if there are multiple words, those are also reversed.
func WordFlip(str string) string {
	slice := strings.Split(str, " ")
	r := []string{}

	for i := 0; i < len(slice); i++ {
		w := strings.Split(slice[i], "")
		rw := []string{}
		for j := len(w) - 1; j >= 0; j-- {
			rw = append(rw, w[j])
		}
		r = append(r, strings.Join(rw, ""))
	}
	return strings.Join(r, " ")
}
