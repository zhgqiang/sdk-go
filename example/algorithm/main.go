package main

import (
	"context"
	"github.com/felix-186/sdk-go/algorithm"
	"github.com/felix-186/sdk-go/example/algorithm/app"
)

func main() {
	s := new(app.TestAlgorithm)
	s.Ctx, s.Cancel = context.WithCancel(context.Background())
	algorithm.NewApp().Start(s)
}
