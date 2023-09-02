package main

import (
	"fmt"
	mypkg "packagework/mypkg"
)

func main() {
	var value mypkg.MyInterface
	value = mypkg.MyType(5)
	value.MethodWithoutParameters()
	value.MethodWithParamaters(127.3)
	fmt.Println(value.MethodWithReturnValue())

}
