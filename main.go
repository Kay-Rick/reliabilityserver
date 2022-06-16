package main

import (
	"rs/consul"
	"rs/healthcheck"
	"time"
)

func main() {
	//每2s检查状态信息
	//创建定时器
	//t := NewMyTick(2, testPrint)
	//t.Start()
	Include(healthcheck.Routers, consul.Routers)
	r := Init()
	reliability := consul.GetReliability()
	if reliability {
		go func() {
			client := consul.ConsulClientInit()
			for {
				consul.StatusCheck(client)
				time.Sleep(4 * time.Second)
			}
		}()
	}

	r.Run(":8000")
}

// 定义函数类型
type Fn func()

// 定时器中的成员
type MyTicker struct {
	MyTick *time.Ticker
	Runner Fn
}

func NewMyTick(interval int, f Fn) *MyTicker {
	return &MyTicker{
		MyTick: time.NewTicker(time.Duration(interval) * time.Second),
		Runner: f,
	}
}

// 启动定时器需要执行的任务
func (t *MyTicker) Start() {
	for {
		select {
		case <-t.MyTick.C:
			t.Runner()
		}
	}
}
