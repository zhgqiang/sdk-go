package main

import (
	"github.com/felix-186/sdk-go/driver"
	"github.com/felix-186/sdk-go/example/driver/app"
)

func main() {
	// 创建采集主程序
	d := new(app.TestDriver)
	driver.NewApp().Start(d)
}
