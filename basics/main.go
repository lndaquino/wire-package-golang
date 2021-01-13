package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/google/wire"
)

type Foo struct {
	X int
}

func ProvideFoo() Foo {
	return Foo{X: 42}
}

type Bar struct {
	X int
}

func ProvideBar(foo Foo) Bar {
	return Bar{X: -foo.X}
}

type Baz struct {
	X int
}

func ProvideBaz(bar Bar) (Baz, error) {
	if bar.X == 0 {
		return Baz{}, errors.New("cannot provide baz when bar is zero")
	}
	return Baz{X: bar.X}, nil
}

var Set = wire.NewSet(ProvideFoo, ProvideBar, ProvideBaz)

func main() {
	e, err := InitializeBaz()
	if err != nil {
		fmt.Printf("failed to create baz: %s\n", err)
		os.Exit(2)
	}
	fmt.Println(e.X)
}
