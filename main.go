package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Set the pyramid's height: ")
	hs := bufio.NewScanner(os.Stdin)
	hs.Scan()
	h := hs.Text()
	hi, _ := strconv.Atoi(h)
	a := make([]int, 0, hi*2)
	length := make([]int, cap(a)+1)
	i := 1
	for index, _ := range length {
		if index <= len(length)/2 {
			a = append(a, index)
			fmt.Println(a)
		} else {
			a := a[0 : index-i]
			fmt.Println(a)
			i = i + 2
			if len(a) == 2 && cap(a)%2 == 1 {
				a := a[0:1]
				fmt.Println(a)
			}

		}
	}
}

