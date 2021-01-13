//+build wireinject

package main

import (
	"github.com/google/wire"
)

func InitializeBaz() (Baz, error) {
	wire.Build(Set)
	return Baz{}, nil
}
