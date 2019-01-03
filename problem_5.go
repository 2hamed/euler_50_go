package main

import "fmt"

func main() {
	for i := 21; true; i++ {
		for j := 20; j >= 3; j-- {
			if i%j != 0 {
				goto cont
			}
		}
		fmt.Println("Product: ", i)
		return
	cont:
		{
			continue
		}
	}

}
