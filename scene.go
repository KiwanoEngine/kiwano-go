package kiwano

import "time"

type Scene interface {
	OnEnter()
	OnExit()
	OnUpdate(time.Duration)
}
