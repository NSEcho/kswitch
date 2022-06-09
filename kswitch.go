package kswitch

import (
	"os"
	"time"
)

// Action is alias for func() and will get executed once ConditionFn return true
type Action func()

// ConditionFn is alias for func() bool; if it returns true, Action func will be executed
type ConditionFn func() bool

// KillSwitch is the main type
type KillSwitch struct {
	interval time.Duration
}

// Returns the new KillSwitch with specific running interval
func NewKillSwitch(interval int) *KillSwitch {
	return &KillSwitch{
		interval: time.Duration(interval),
	}
}

// Run will execute cf() every interval seconds and once cf() return true, ac() will be called and the program will exit
func (k *KillSwitch) Run(cf ConditionFn, ac Action) {
	t := time.NewTicker(k.interval * time.Second)

	for {
		<-t.C
		if cf() {
			ac()
			os.Exit(1)
		}
	}
}
