package core

import (
	"fmt"
	"testing"
	"time"
)

func TestGetTimeWheel(t *testing.T) {
	timeWheel := GetTimeWheel(10,time.Second)
	timeWheel.AddTask("world",func(args interface{}){
		fmt.Println("hello",args)
	},time.Second*1)

	time.Sleep(time.Second*2)
}


func TestGetSlotCount(t *testing.T){
	timeWheel := GetTimeWheel(10,time.Second)
	tests := []struct {
		name   string
		args time.Duration
		want int
	}{
		{
			name: "case1: one min",
			args: time.Minute,
			want: 60,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := timeWheel.getSlotCount(tt.args); got != tt.want {
				t.Errorf("TestGetSlotCount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetloopCount(t *testing.T){
	timeWheel := GetTimeWheel(10,time.Second)
	tests := []struct {
		name   string
		args int
		want int
	}{
		{
			name: "case1: one loop",
			args: 12,
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := timeWheel.getLoopCount(tt.args); got != tt.want {
				t.Errorf("TestGetloopCount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetTaskIndex(t *testing.T) {
	timeWheel := GetTimeWheel(10,time.Second)
	tests := []struct {
		name   string
		args int
		want int
	}{
		{
			name: "case1: one times",
			args: 60,
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := timeWheel.getTaskIndex(tt.args); got != tt.want {
				t.Errorf("TestGetTaskIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddTask(t *testing.T){
	timeWheel := GetTimeWheel(10,time.Second)
	tests := []struct {
		name   string
		args interface{}
		wheelSlotFunc WheelSlotFunc
		delayTime time.Duration
		wantTaskCount int
	}{
		{
			name: "case1: one times",
			args: "world",
			wheelSlotFunc: func(args interface{}){
				fmt.Println("hello",args)
			},
			delayTime:time.Second,
			wantTaskCount: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			timeWheel.AddTask(tt.args,tt.wheelSlotFunc,tt.delayTime)
			gotCount := len(timeWheel.wheelSlots[1])
			if gotCount != tt.wantTaskCount {
				t.Errorf("TestGetTaskIndex() = %v, want %v", gotCount, tt.wantTaskCount)
			}

		})
	}
}