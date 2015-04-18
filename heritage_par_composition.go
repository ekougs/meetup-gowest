package main

import "fmt"

type Bar struct {
	bar string
}

type Foo struct {
	Bar // Anonymous member
	foo string
}

func main() {
	foo := Foo{}
	foo.Bar.bar = "Bar"
	foo.foo = "foo"
	fmt.Println(foo)
	// Pas besoin de facade pour utiliser directement bar
	foo.bar = "bar"
	fmt.Println(foo)
}

