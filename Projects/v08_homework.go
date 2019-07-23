package main

import (
	"fmt"
)

func main() {
	m1 := map[int]string{1: "a", 2: "b", 3: "c", 4: "d", 5: "e"}
	m2 := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5}
	fmt.Println(m1) // map[1:a 2:b 3:c 4:d 5:e]
	fmt.Println(m2) // map[a:1 b:2 c:3 d:4 e:5]

	m3 := make(map[string]int, 5)
	for k, v := range m1 {
		m3[v] = k
	}
	fmt.Println(m3) // map[a:1 b:2 c:3 d:4 e:5]
}
