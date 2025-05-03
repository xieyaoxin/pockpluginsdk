package runner

import (
	"context"
	"fmt"
	"github.com/looplab/fsm"
	"time"
)

// 使用状态机重写调度方法
var STATE_MACHINE *fsm.FSM
var ctx = context.Background()

const (
	CLOSED          = "CLOSED"
	BATTLE_RUNNING  = "BATTLE_RUNNING"
	MERGE_RUNNING   = "MERGE_RUNNING"
	NIRVANA_RUNNING = "NIRVANA_RUNNING"
	TT_RUNNING      = "TT_RUNNING"
	DUNGEON_RUNNING = "DUNGEON_RUNNING"
)

const (
	CLOSE   = "CLOSE"
	BATTLE  = "BATTLE"
	MERGE   = "MERGE"
	NIRVANA = "NIRVANA"
	TT      = "TT"
	DUNGEON = "DUNGEON"
)

var HandlerMap = map[string]StateMachineSpi{}

func init() {
	STATE_MACHINE = fsm.NewFSM(
		CLOSED,
		fsm.Events{
			{Name: BATTLE, Src: []string{CLOSED}, Dst: BATTLE_RUNNING},
			{Name: MERGE, Src: []string{CLOSED}, Dst: MERGE_RUNNING},
			{Name: NIRVANA, Src: []string{CLOSED}, Dst: NIRVANA_RUNNING},
			{Name: TT, Src: []string{CLOSED}, Dst: TT_RUNNING},
			{Name: DUNGEON, Src: []string{CLOSED}, Dst: DUNGEON_RUNNING},
			{Name: CLOSE, Src: []string{BATTLE_RUNNING, MERGE_RUNNING, NIRVANA_RUNNING, TT_RUNNING, DUNGEON_RUNNING}, Dst: CLOSED},
		},
		fsm.Callbacks{
			"before_event": func(ctx context.Context, e *fsm.Event) {
				fn := HandlerMap[e.Src]
				if fn != nil {
					fn.HandleFinishEvent()
					time.Sleep(1 * time.Second)
				}
			},
			"leave_state": func(ctx context.Context, e *fsm.Event) {
				//	 确认任务已完成
				fmt.Printf("leave_state, ctx is %v ,event is %v  \n", ctx, e)
			},
			"enter_state": func(ctx context.Context, e *fsm.Event) {
				fn := HandlerMap[e.Dst]
				if fn != nil {
					fn.HandleStartEvent()
					print("开启任务成功")
				}
			},
			"after_event": func(ctx context.Context, e *fsm.Event) {
				//	 确认任务已完成
				fmt.Printf("after_event, ctx is %v ,event is %v  \n", ctx, e)

			},
		},
	)
}

func StopTask() error {
	return STATE_MACHINE.Event(ctx, CLOSE)
}

func StartTask(taskName string) error {
	return STATE_MACHINE.Event(ctx, taskName)
}
