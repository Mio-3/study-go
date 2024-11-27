package main

import (
	"fmt"
)

type MyInt int
func (yu *MyInt) Inc() { *yu++ }

func change() {
	// var sum int
	// sum = 5 + 6 + 3
	// avg := float64(sum) / 3
	// if avg > 4.5 {
	// 	fmt.Println("good")
	// }

	// ns := [...]int{1, 2, 3, 4, 5}
	// fmt.Println(ns[3])

	// fmt.Println(len(ns))

	// fmt.Println(ns[1:3])

	// ns1 := []int{10, 20, 30, 40, 50}

	// fmt.Println(ns1[3])

	// fmt.Println(len(ns1))

	// fmt.Println(cap(ns1))

	// ns1 = append(ns1, 60, 70)
	// fmt.Println(len(ns1), cap(ns1))

	// n1 := 19
	// n2 := 86
	// n3 := 1
	// n4 := 12

	// sum := n1 + n2 + n3 + n4
	// fmt.Println(sum)
	ns1 := []int{19, 86, 1, 12}
	var sum int
	for i := 0; i < len(ns1); i++ {
		sum += ns1[i]
	}
	fmt.Println(sum)

	m := map[string]int{"x": 10, "y": 20}

	fmt.Println(m["x"])

	m["z"] = 30
	n, ok := m["z"]
	fmt.Println(n, ok) 

	delete(m, "z")

	n, ok = m["z"]
	fmt.Println(n, ok)

	type Score struct {
		UserID string
		GameID int
		Point int
	}

	fs := make([]func() string, 2)
	fs[0] = func() string { return "hoge" }
	fs[1] = func() string { return "fuga" }

	for _, f := range fs {
		fmt.Println(f())
	}

	for i := 1; i <= 100; i++ {
		print(i)
		if i % 2 == 0 {
			fmt.Println("-偶数")
		} else {
			fmt.Println("-奇数")
		}
	}
	n1, m1 := 10, 20

	swap2(&n1, &m1)
	fmt.Println(n1, m1)

	var yu MyInt
	fmt.Println(yu)
	(&yu).Inc()
	fmt.Println(yu)

}

func add(x, y int) (int, int) {
	return y, x
}

func swap(x, y int) (int, int) {
	return y, x
}

func swap2(x, y *int) {
	*x, *y = *y, *x
}