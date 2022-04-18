package common

type State string

type StateChanger interface {
	SetState(state State)
	GetState() State
}
