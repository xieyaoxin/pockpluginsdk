package runner

type StateMachineSpi interface {
	HandleFinishEvent()
	HandleStartEvent()
	HandleAfterEvent()
}
