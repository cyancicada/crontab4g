package crontab4g

import "time"

// ConstantDelaySchedule表示一个简单的循环占空比，例如 “每5分钟一次”。
//它不支持比一秒钟更频繁的工作。
type ConstantDelaySchedule struct {
	Delay time.Duration
}

//每个返回一个crontab计划，每个持续时间激活一次。
//不支持小于一秒的延迟（最多可延迟1秒）。
//截断任何小于一秒的字段。
func Every(duration time.Duration) ConstantDelaySchedule {
	if duration < time.Second {
		duration = time.Second
	}
	return ConstantDelaySchedule{
		Delay: duration - time.Duration(duration.Nanoseconds())%time.Second,
	}
}

// Next将在下次运行时返回。
//此轮次以便下一个激活时间将在第二个激活时间
func (schedule ConstantDelaySchedule) Next(t time.Time) time.Time {
	return t.Add(schedule.Delay - time.Duration(t.Nanosecond())*time.Nanosecond)
}
