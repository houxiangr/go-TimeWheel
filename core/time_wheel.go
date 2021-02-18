package core

import (
	"sync"
	"time"
)

const MinIntervalTime = 1

type WheelSlotFunc func(args interface{}) error

type slotEntity struct {
	args          interface{}
	wheelSlotFunc WheelSlotFunc
	//exec be left cycle count
	loopCount int
}

func (this *slotEntity) Init(args interface{}, wheelSlotFunc WheelSlotFunc, loopCount int) {
	this.loopCount = loopCount
	this.args = args
	this.wheelSlotFunc = wheelSlotFunc
}

type wheetSlot []slotEntity

type TimeWheel struct {
	slotSize int
	//min 1ms
	slotTimeInterval time.Duration

	wheelSlots  []wheetSlot
	currentSlot int

	// this time wheel task count
	// if task count is zero,suspend the time wheel
	taskCount int

	//fail task slice
	failTask []slotEntity

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
	this.failTask = []slotEntity{}
	this.lock = sync.RWMutex{}
}

func (this *TimeWheel) Start() {
	for {
		if this.taskCount == 0 {
			this.stopLoop()
			break
		}
		//fmt.Println("current slot",this.currentSlot)
		this.lock.RLock()
		defer this.lock.RUnlock()
		//exec current slot function
		slotEntity := this.wheelSlots[this.currentSlot]
		entityLen := len(slotEntity)
		for i := 0; i < entityLen; i++ {
			if slotEntity[i].loopCount > 0 {
				slotEntity[i].loopCount--
				continue
			}
			err := slotEntity[i].wheelSlotFunc(slotEntity[i].args)
			// insert fail task
			if err != nil {
				this.dealFailTask(slotEntity[i])
				continue
			}
			if i == entityLen-1 {
				slotEntity = slotEntity[:i]
			} else {
				slotEntity = append(slotEntity[:i], slotEntity[i+1:]...)
				i--
				entityLen--
				this.taskCount--
			}
		}

		//update current slot index
		this.currentSlot = (this.currentSlot + 1) % this.slotSize

		//sleep
		time.Sleep(this.slotTimeInterval)
	}
}

func (this *TimeWheel) AddTask(args interface{}, wheelSlotFunc WheelSlotFunc, delayTime time.Duration) {
	slotCount := this.getSlotCount(delayTime)
	//structure the insert slot
	slotEntity := slotEntity{}
	slotEntity.Init(args, wheelSlotFunc, this.getLoopCount(slotCount))

	this.lock.Lock()
	defer this.lock.Unlock()
	//insert right slot
	this.taskCount++
	this.wheelSlots[this.getTaskIndex(slotCount)] = append(this.wheelSlots[this.getTaskIndex(slotCount)], slotEntity)
	// if task count == 1,restart loop
	if this.taskCount == 1 {
		go this.Start()
	}
}

func (this *TimeWheel) getLoopCount(slotCount int) int {
	return slotCount / this.slotSize
}

//calc the wheel cycle times
func (this *TimeWheel) getSlotCount(delayTime time.Duration) int {
	return int(delayTime.Milliseconds() / this.slotTimeInterval.Milliseconds())
}

//calc the task right index
func (this *TimeWheel) getTaskIndex(loopCount int) int {
	return (this.currentSlot + loopCount%this.slotSize) % this.slotSize
}

//calc the task right index
func (this *TimeWheel) dealFailTask(failTask slotEntity) {
	this.failTask = append(this.failTask, failTask)
}

//calc the task right index
func (this *TimeWheel) stopLoop() {
	this.currentSlot = 0
}

func GetTimeWheel(slotSize int, slotTimeInterval time.Duration) TimeWheel {
	timeWheel := TimeWheel{}
	timeWheel.Init(slotSize, slotTimeInterval)
	go timeWheel.Start()

	return timeWheel
}
