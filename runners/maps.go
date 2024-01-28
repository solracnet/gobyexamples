package runners

import (
	"fmt"
	"maps"
)

type User struct {
	ID   int32
	Name string
}

func Maps() {

	u := User{ID: 1, Name: "Carlos"}

	m := make(map[string]any)

	m["k1"] = 7
	m["k2"] = 13
	m["315478"] = u

	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("v1:", v1)

	v3 := m["k3"]
	fmt.Println("v3:", v3)
	fmt.Println("v3 exists:", v3 != nil)
	fmt.Println("k3 exists:", m["k3"] != nil)

	fmt.Println("len:", len(m))

	delete(m, "k2")
	fmt.Println("map:", m)

	clear(m)
	fmt.Println("map after clear:", m)

	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
	fmt.Println("map len:", len(n))

	n2 := map[string]int{"foo": 1, "bar": 2}
	if maps.Equal(n, n2) {
		fmt.Println("n == n2")
	}
}
