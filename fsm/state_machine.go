package fsm

import (
	"fmt"
	"github.com/looplab/fsm"
)

func RefStockStateMachine() *StockStateMachine {
	return stateMachine
}

type StockStateMachine struct {
	Fsm *fsm.FSM
}

var stateMachine = &StockStateMachine{}

func Init() {
	stateMachine.Fsm = fsm.NewFSM(
		"sightseer",
		fsm.Events{

			{Name: "openAccount", Src: []string{"sightseer"}, Dst: "buyer"},

			{Name: "buy", Src: []string{"buyer"}, Dst: "holder"},

			{Name: "sellPartial", Src: []string{"holder"}, Dst: "holder"},

			{Name: "sellAll", Src: []string{"holder"}, Dst: "buyer"},

			{Name: "writeOff", Src: []string{"buyer"}, Dst: "sightseer"},
		},
		fsm.Callbacks{
			"leave_state": func(event *fsm.Event) { fmt.Printf("now leaving state:%v to %v \n", event.Src, event.Dst) },

			"leave_sightseer": ExitsightseerFunc,
			"enter_sightseer": EntersightseerFunc,

			"leave_buyer": ExitbuyerFunc,
			"enter_buyer": EnterbuyerFunc,

			"leave_holder": ExitholderFunc,
			"enter_holder": EnterholderFunc,
		},
	)
}
