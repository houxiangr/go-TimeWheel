package core

import "time"

/*
depends on delayTime auto create suitable size time wheel。
 */

/*
size class
depend on delayTime
 */
type TimeWheelSize int
const (
	TimeWheelSizeMillisecond TimeWheelSize = iota //1ms <= delaytime < 1000ms slotsize = 100 slotTimeInterval = 1ms
	TimeWheelSizeSecond //delayTime not to be divisible by one min,slotsize =100 slotTimeInterval = 1s
	TimeWheelSizeMinute //delayTime to be divisible by one min slotsize =100 slotTimeInterval = 1min
	TimeWheelSizeHour //delayTime to be divisible by one hour slotsize =100 slotTimeInterval = 1hour
	TimeWheelSizeDay //delayTime to be divisible by one day slotsize =100 slotTimeInterval = 1day

	DefaultSlotSize = 100
)

type TimeWheelMap map[TimeWheelSize]*TimeWheel

func (this *TimeWheelMap) AddTask(args interface{}, wheelSlotFunc WheelSlotFunc,delayTime time.Duration){
	timeWheelSize := calcDelayTimeBelongClass(delayTime)
	targetTimeWheel,ok := (*this)[timeWheelSize]
	if !ok || targetTimeWheel == nil {
		tempTimeWheel := GetTimeWheel(DefaultSlotSize,getSlotTimeIntervalByTimeWheelSize(timeWheelSize))
		tempTimeWheel.AddTask(args,wheelSlotFunc,delayTime)
		(*this)[timeWheelSize] = &tempTimeWheel
		return
	}
	targetTimeWheel.AddTask(args,wheelSlotFunc,delayTime)
}

func GetTimeWheels() TimeWheelMap {
	return TimeWheelMap{}
}

func calcDelayTimeBelongClass(delayTime time.Duration) TimeWheelSize {
	if delayTime.Milliseconds()%(time.Hour*24).Milliseconds() == 0 {
		return TimeWheelSizeDay
	}else if delayTime.Milliseconds()%time.Hour.Milliseconds() == 0 {
		return TimeWheelSizeHour
	}else if delayTime.Milliseconds()%time.Minute.Milliseconds() == 0 {
		return TimeWheelSizeMinute
	}else if delayTime.Milliseconds()%time.Second.Milliseconds() == 0 {
		return TimeWheelSizeSecond
	}else if delayTime.Milliseconds()%time.Millisecond.Milliseconds() == 0 {
		return TimeWheelSizeMillisecond
	}
	return TimeWheelSizeMillisecond
}

func getSlotTimeIntervalByTimeWheelSize(timeWheelSize TimeWheelSize) time.Duration {
	switch timeWheelSize {
	case TimeWheelSizeDay:
		return time.Hour * 24
	case TimeWheelSizeHour:
		return time.Hour
	case TimeWheelSizeMinute:
		return time.Minute
	case TimeWheelSizeSecond:
		return time.Second
	case TimeWheelSizeMillisecond:
		return time.Millisecond
	default:
		return time.Millisecond
	}
}