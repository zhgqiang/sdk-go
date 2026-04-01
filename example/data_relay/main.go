package main

import (
	"context"
	"github.com/felix-186/sdk-go/data_relay"
	"github.com/felix-186/sdk-go/example/data_relay/app"
)

func main() {
	d := new(app.TestRelay)
	d.Ctx, d.Cancel = context.WithCancel(context.Background())
	data_relay.NewApp().Start(d)
}
