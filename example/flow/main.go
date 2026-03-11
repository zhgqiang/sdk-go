package main

import (
	"context"
	"github.com/zhgqiang/sdk-go/example/flow/app"
	"github.com/zhgqiang/sdk-go/flow"
)

func main() {
	// 创建采集主程序
	d := new(app.TestFlow)
	d.Ctx, d.Cancel = context.WithCancel(context.Background())
	flow.NewApp().Start(d)
}
