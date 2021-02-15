package core

import (
	"fmt"
	"sync"
	"time"
)

const MinIntervalTime = 1

type WheelSlotFunc func(args interface{})

type slotEntity struct {
	args          interface{}
	wheelSlotFunc WheelSlotFunc
}

type wheetSlot []slotEntity

type TimeWheel struct {
	slotSize         int
	//min 1ms
	slotTimeInterval time.Duration

	wheelSlots  []wheetSlot
	currentSlot int

	lock sync.RWMutex
}

func (this TimeWheel) Init(slotSize int, slotTimeInterval time.Duration) {
	this.slotSize = slotSize
	if slotTimeInterval.Milliseconds() < MinIntervalTime {
		slotTimeInterval = time.Millisecond
	}
	this.slotTimeInterval = slotTimeInterval
	this.currentSlot = 0
	this.wheelSlots = make([]wheetSlot, slotSize)
	this.lock = sync.RWMutex{}
}

func (this TimeWheel) Start(){
	for {
		fmt.Println("test")
	}
}

func (this TimeWheel) AddTask(args interface{}, wheelSlotFunc WheelSlotFunc,delayTime time.Duration) {

}

func GetTimeWheel(slotSize int, slotTimeInterval time.Duration) TimeWheel {
	timeWheel := TimeWheel{}
	timeWheel.Init(slotSize, slotTimeInterval)

	return timeWheel
}
