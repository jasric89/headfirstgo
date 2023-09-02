package mypkg

import "fmt"

type MyInterface interface {
	MethodWithoutParameters()
	MethodWithParamaters(float64)
	MethodWithReturnValue() string
}

type MyType int

func (m MyType) MethodWithoutParameters() {
	fmt.Println("MethodWithoutParameters Called")
}

func (m MyType) MethodWithParamaters(f float64) {
	fmt.Println("MethodWithParamaters Called", f)
}

func (m MyType) MethodWithReturnValue() string {
	return "MethodWithReturnValue Called"
}
