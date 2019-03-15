package main

// TODO: 为啥只能导入 module 目录， 不能只导入module下的某个文件?
import "./module"

func main() {
	module.Test1()
	module.Test2()
}
