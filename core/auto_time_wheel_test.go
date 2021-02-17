package core

import (
	"fmt"
	"testing"
	"time"
)

func TestGetTimeWheels(t *testing.T) {
	timeWheels := GetTimeWheels()
	timeWheels.AddTask("world",func(args interface{}){
		fmt.Println("hello",args)
	},time.Second*5)

	time.Sleep(time.Second*10000)
}

func TestAutoAddTask(t *testing.T){
	timeWheels := GetTimeWheels()
	timeWheels.AddTask("world",func(args interface{}){
		fmt.Println("hello",args)
	},time.Second*5)
	if timeWheels[TimeWheelSizeSecond] == nil {
		t.Error("auto create wheel fail")
	}

	timeWheels.AddTask("world",func(args interface{}){
		fmt.Println("hello",args)
	},time.Second*5)
	if len(timeWheels[TimeWheelSizeSecond].wheelSlots[5]) != 2 {
		t.Error("repeat auto create wheel fail")
	}
}

func TestCalcDelayTimeBelongClass(t *testing.T){
	tests := []struct {
		name   string
		args time.Duration
		want TimeWheelSize
	}{
		{
			name: "delay time is one hour",
			args: time.Hour*24,
			want: TimeWheelSizeDay,
		},
		{
			name: "delay time is one hour",
			args: time.Hour,
			want: TimeWheelSizeHour,
		},
		{
			name: "delay time is one hour",
			args: time.Minute,
			want: TimeWheelSizeMinute,
		},
		{
			name: "delay time is one hour",
			args: time.Second,
			want: TimeWheelSizeSecond,
		},
		{
			name: "delay time is one hour",
			args: time.Millisecond,
			want: TimeWheelSizeMillisecond,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calcDelayTimeBelongClass(tt.args); got != tt.want {
				t.Errorf("TestCalcDelayTimeBelongClass() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetSlotTimeIntervalByTimeWheelSize(t *testing.T){
	tests := []struct {
		name   string
		args TimeWheelSize
		want time.Duration
	}{
		{
			name: "delay time is one hour",
			want: time.Hour*24,
			args: TimeWheelSizeDay,
		},
		{
			name: "delay time is one hour",
			want: time.Hour,
			args: TimeWheelSizeHour,
		},
		{
			name: "delay time is one hour",
			want: time.Minute,
			args: TimeWheelSizeMinute,
		},
		{
			name: "delay time is one hour",
			want: time.Second,
			args: TimeWheelSizeSecond,
		},
		{
			name: "delay time is one hour",
			want: time.Millisecond,
			args: TimeWheelSizeMillisecond,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getSlotTimeIntervalByTimeWheelSize(tt.args); got != tt.want {
				t.Errorf("TestCalcDelayTimeBelongClass() = %v, want %v", got, tt.want)
			}
		})
	}
}
