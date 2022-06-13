package main

import "fmt"

func main() {

	for i := 3; i <= 100; i++ {
		//b := i
		b := 0

		for j := 2; j < i; j++ {

			if j != i {
				if i%j == 0 {
					b = 1

				} else {

				}

			}

		}

		if b == 1 {

		} else {
			fmt.Println(i)
		}

	}

}
