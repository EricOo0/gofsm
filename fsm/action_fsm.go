package fsm

import (
	"fmt"
	"github.com/looplab/fsm"
)

func EnterStat1Func(event *fsm.Event) {
	fmt.Println("an event is gonna happen")
}
