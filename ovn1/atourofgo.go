package main

import (
	"fmt"
	"golang.org/x/tour/pic"
	"golang.org/x/tour/wc"
	"math"
	"strings"
)

func Sqrt(x float64) float64 {
	z := float64(1)
	for {
		old := z
		z -= (z*z - x) / (2 * z)
		fmt.Println(z)
		if old-z < 0.000000000001 && old-z > 0 {
			break
		}
	}
	return z
}

func Pic(dx, dy int) [][]uint8 {
	arr := make([][]uint8, dy)
	for i := range arr {
		arr[i] = make([]uint8, dx)
		for j := range arr[i] {
			arr[i][j] = uint8(i * j)
		}
	}
	return arr
}

func WordCount(s string) map[string]int {
	w := strings.Fields(s)
	m := make(map[string]int)
	for _, e := range w {
		m[e]++
	}
	return m
}

func fibonacci() func() int {
	x, y := 0, 0
	return func() int {
		x, y = y, x+y
		if y == 0 {
			y = 1
		}
		return x
	}
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(math.Sqrt(2))
	pic.Show(Pic)
	wc.Test(WordCount)
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
