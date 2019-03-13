package main

import "fmt"

func main()  {
	const s = "Go语言"

	for i, r := range s {
		fmt.Printf("%#U  ： %d\n", r, i)
	}
	/*
	U+0047 'G'  ： 0
	U+006F 'o'  ： 1
	U+8BED '语'  ： 2
	U+8A00 '言'  ： 5	
	*/
}
