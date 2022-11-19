package fsm

//
//import (
//	"fmt"
//	"github.com/looplab/fsm"
//)
//
//func RefStockStateMachine() *StockStateMachine {
//	return stockStateMachine
//}
//
//type StockStateMachine struct {
//	fsm *fsm.FSM
//}
//
//var stockStateMachine = &StockStateMachine{}
//
//func Init() {
//	stockStateMachine.fsm = fsm.NewFSM(
//		"sightseer",
//		fsm.Events{
//			{Name: "openAccount", Src: []string{"sightseer"}, Dst: "buyer"},
//			{Name: "buy", Src: []string{"buyer"}, Dst: "holder"},
//			{Name: "sellPartial", Src: []string{"holder"}, Dst: "holder"},
//			{Name: "sellAll", Src: []string{"holder"}, Dst: "buyer"},
//			{Name: "writeOff", Src: []string{"buyer"}, Dst: "sightseer"},
//		},
//		fsm.Callbacks{
//			"before_event":       func(event *fsm.Event) { fmt.Println("an event is gonna happen") },
//			"before_openAccount": func(event *fsm.Event) { fmt.Println("event 'openAccount' is gonna happen") },
//			"before_buy":         func(event *fsm.Event) { fmt.Println("event 'buy' is gonna happen") },
//			"before_sellPartial": func(event *fsm.Event) { fmt.Println("event 'sellPartial' is gonna happen") },
//			"before_sellAll":     func(event *fsm.Event) { fmt.Println("event 'sellAll' is gonna happen") },
//			"before_writeOff":    func(event *fsm.Event) { fmt.Println("event 'write off' is gonna happen") },
//			"leave_state":        func(event *fsm.Event) { fmt.Printf("now leaving state:%v to %v \n", event.Src, event.Dst) },
//		},
//	)
//}
