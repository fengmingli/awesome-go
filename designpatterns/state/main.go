package main

/**
 * @Author: LFM
 * @Date: 2023/5/31 23:41
 * @Since: 1.0.0
 * @Desc: TODO
 */

import "fmt"

// State 状态枚举
type State int

const (
	StateA State = iota
	StateB
	StateC
)

// Event 事件枚举
type Event int

const (
	Event1 Event = iota
	Event2
	Event3
)

// StateMachine 状态机结构体
type StateMachine struct {
	currentState State
}

// 状态转移函数
func (sm *StateMachine) transition(event Event) {
	switch sm.currentState {
	case StateA:
		if event == Event1 {
			fmt.Println("Transition from StateA to StateB")
			sm.currentState = StateB
		}
	case StateB:
		if event == Event2 {
			fmt.Println("Transition from StateB to StateC")
			sm.currentState = StateC
		}
	case StateC:
		if event == Event3 {
			fmt.Println("Transition from StateC to StateA")
			sm.currentState = StateA
		}
	}
}

// 主函数
func main() {
	sm := StateMachine{currentState: StateA}

	// 执行一系列事件，触发状态转移
	sm.transition(Event1)
	sm.transition(Event2)
	sm.transition(Event3)
	sm.transition(Event1)
}
