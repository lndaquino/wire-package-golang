//+build wireinject

package main

import (
	"github.com/google/wire"
)

func InitializeBar() (string, error) {
	panic(wire.Build(Set))
}
