package main

import "fmt"

// SumInts adds together the values of m.
func SumInts(m map[string]int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}

	return s
}

// SumFloats adds together the values of m.
func SumFloats(m map[string]float64) float64 {
	var s float64
	for _, v := range m {
		s += v
	}

	return s
}

// SumIntsOrFloats sums the values of map m. It supports both Int64 and float64
// as types for map values.

// Go requires that map keys be comparable
func SumIntsOrFloats[k comparable, V int64 | float64](m map[k]V) V {
	var s V
	for _, v := range m {
		s += v
	}

	return s
}

func SumIntsOrFloatsOrString[k comparable, V int64 | float64 | string](m map[k]V) V {
	var s V
	for _, v := range m {
		s += v
	}

	return s
}

func main() {
	ints := map[string]int64{
		"first":  34,
		"second": 12,
	}
	floats := map[string]float64{
		"first":  35.98,
		"second": 26.99,
	}

	fmt.Printf("Non-generic Sums: %v and %v\n",
		SumInts(ints),
		SumFloats(floats),
	)

	fmt.Printf("Generic Sums: %v and %v\n",
		SumIntsOrFloats[string, int64](ints),
		SumIntsOrFloats[string, float64](floats),
	)

	// We also ommit type specifications
	fmt.Printf("Generics Sums wo type specifications: %v and %v\n",
		SumIntsOrFloats(ints),
		SumIntsOrFloats(floats),
	)

	// Because K is comparable, we can also use integer
	intsWithInt := map[int]int64{
		1: 10,
		2: 20,
	}
	fmt.Println(SumIntsOrFloats[int, int64](intsWithInt))

	strs := map[int]string{
		1: "hoge",
		2: "fuga",
	}
	fmt.Println(SumIntsOrFloatsOrString(strs))
}
