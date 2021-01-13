package main

import (
	"fmt"
	"os"

	"github.com/google/wire"
)

type Fooer interface {
	Foo() string
}

type MyFooer string

func (b *MyFooer) Foo() string {
	return string(*b)
}

func provideMyFooer() *MyFooer {
	b := new(MyFooer)
	*b = "Hello, World!"
	return b
}

type Bar string

func provideBar(f Fooer) string {
	// f will be a *MyFooer.
	return f.Foo()
}

var Set = wire.NewSet(
	provideMyFooer,
	wire.Bind(new(Fooer), new(*MyFooer)),
	provideBar)

func main() {
	e, err := InitializeBar()
	if err != nil {
		fmt.Printf("failed to create bar: %s\n", err)
		os.Exit(2)
	}
	fmt.Println(e)
}
