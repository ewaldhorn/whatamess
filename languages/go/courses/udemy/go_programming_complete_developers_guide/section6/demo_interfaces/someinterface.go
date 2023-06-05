package main

type MyInterface interface {
	Function1()
	Function2(x int) int
}

type MyType int

func (m MyType) Function1() {}

func (m MyType) Function2(x int) int {
	return x + x
}

func execute(i MyInterface) {
	i.Function1()
}
