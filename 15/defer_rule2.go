// 规则二 defer执行顺序为先进后出

package main

import "fmt"

func main() {
	defer fmt.Print(" !!! ")
	defer fmt.Print(" world ")
	fmt.Print(" hello ")

}

//输出:  hello  world  !!!
