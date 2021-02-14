package core

import (
	"time"
)

type WheelSlotFunc func(args interface{})

type slotEntity struct {
	slotId        string
	args          interface{}
	wheelSlotFunc WheelSlotFunc
}

type wheetSlot []slotEntity

type TimeWheel struct {
	slotSize         int
	slotTimeInterval time.Duration

	wheelSlots  []wheetSlot
	currentSlot int
}

func (this TimeWheel) InitTimeWheel(slotSize int, slotTimeInterval time.Duration) {
	this.slotSize = slotSize
	this.slotTimeInterval = slotTimeInterval
	this.currentSlot = 0
	this.wheelSlots = make([]wheetSlot, slotSize)
}

func (this TimeWheel) AddTask(args interface{}, wheelSlotFunc WheelSlotFunc,delayTime time.Duration) {

}

func GetTimeWheel(slotSize int, slotTimeInterval time.Duration) TimeWheel {
	timeWheel := TimeWheel{}
	timeWheel.InitTimeWheel(slotSize, slotTimeInterval)
	return timeWheel
}
