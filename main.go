package main

import (
	"github.com/go-TimeWheel/core"
	"time"
)

func main(){
	timeWheel := core.GetTimeWheel(10,time.Second)
	timeWheel.AddTask(1,func(args interface{}){},time.Second)
}
