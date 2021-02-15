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
	//exec be left cycle count
	loopCount int
}

func (this *slotEntity)Init(args interface{},wheelSlotFunc WheelSlotFunc,loopCount int){
	this.loopCount = loopCount
	this.args = args
	this.wheelSlotFunc = wheelSlotFunc
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

func (this *TimeWheel) Init(slotSize int, slotTimeInterval time.Duration) {
	this.slotSize = slotSize
	if slotTimeInterval.Milliseconds() < MinIntervalTime {
		slotTimeInterval = time.Millisecond
	}
	this.slotTimeInterval = slotTimeInterval
	this.currentSlot = 0
	this.wheelSlots = make([]wheetSlot, slotSize)
	this.lock = sync.RWMutex{}
}

func (this *TimeWheel) Start(){
	for {
		fmt.Println("current slot",this.currentSlot)
		this.lock.RLock()
		defer this.lock.RUnlock()
		//exec current slot function
		slotEntity := this.wheelSlots[this.currentSlot]
		entityLen := len(slotEntity)
		for i:=0;i<entityLen;i++ {
			if slotEntity[i].loopCount > 0 {
				slotEntity[i].loopCount--
				continue
			}
			slotEntity[i].wheelSlotFunc(slotEntity[i].args)
			if i == entityLen-1 {
				slotEntity = slotEntity[:i]
			} else {
				slotEntity = append(slotEntity[:i], slotEntity[i+1:]...)
				i--
				entityLen--
			}
		}

		//update current slot index
		this.currentSlot = (this.currentSlot+1)%this.slotSize

		//sleep
		time.Sleep(this.slotTimeInterval)
	}
}

func (this *TimeWheel) AddTask(args interface{}, wheelSlotFunc WheelSlotFunc,delayTime time.Duration) {
	loopCount := int(delayTime.Milliseconds()/this.slotTimeInterval.Milliseconds())
	//structure the insert slot
	slotEntity := slotEntity{}
	slotEntity.Init(args,wheelSlotFunc,loopCount/this.slotSize)

	this.lock.Lock()
	defer this.lock.Unlock()
	//insert right slot
	this.wheelSlots[(this.currentSlot + loopCount%this.slotSize)%this.slotSize] = append(this.wheelSlots[(this.currentSlot + loopCount%this.slotSize)%this.slotSize],slotEntity)
}

func GetTimeWheel(slotSize int, slotTimeInterval time.Duration) TimeWheel {
	timeWheel := TimeWheel{}
	timeWheel.Init(slotSize, slotTimeInterval)

	return timeWheel
}
