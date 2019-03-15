package main

func main() {
	if a := 1; false {
	} else if b := 2; false {
	} else if c := 3; false {
	} else {
		// 作用域内
		println(a, b, c)
	}

	//  作用域外
	//println(a)
}
