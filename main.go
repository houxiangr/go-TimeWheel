package main

import (
	"fmt"
	"github.com/go-TimeWheel/core"
	"time"
)

func main() {
	timeWheel := core.GetTimeWheel(10, time.Second)
	timeWheel.AddTask(1, func(args interface{}) error {
		fmt.Println("test", args)
		return nil
	}, time.Second*5)

	//timeWheels := core.GetTimeWheels()
	//timeWheels.AddTask(1,func(args interface{})error{
	//		fmt.Println("test",args)
	//		return nil
	//	},time.Second*5)
	//
	time.Sleep(time.Second * 10000)
}
