//go:build wireinject
// +build wireinject

package main

import (
	"github.com/alindesign/snippy/internal"
	"github.com/google/wire"
)

func InitializeApp() (*internal.App, error) {
	wire.Build(
		internal.Set,
	)

	return &internal.App{}, nil
}
