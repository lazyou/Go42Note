package main

import (
	"fmt"
)

func main() {
	s := "Go语言四十二章经"
	for k, v := range s {
		fmt.Printf("k：%d,v：%c == %d\n", k, v, v)
	}
}

/*
k：0,v：G == 71
k：1,v：o == 111
k：2,v：语 == 35821
k：5,v：言 == 35328
k：8,v：四 == 22235
k：11,v：十 == 21313
k：14,v：二 == 20108
k：17,v：章 == 31456
k：20,v：经 == 32463
*/
