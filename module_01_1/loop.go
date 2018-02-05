package main

import (
	"errors"
	"fmt"
)

const len = 100

func main() {
	var arr [len]int

	for i := 0; i < len; i++ {
		var v int
		var err error
		if i%2 == 0 {
			v, err = double(i)
			if err != nil {
				panic(err)
			}
		} else {
			v = i
		}
		arr[i] = v
	}

	for i := 0; i < len; i++ {
		fmt.Printf("%d\n", arr[i])
	}
}

func double(i int) (int, error) {
	if i < 0 {
		return i, errors.New("int cannot be less than 0")
	}
	return (i * 2), nil
}
