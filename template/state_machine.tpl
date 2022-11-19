package fsm

import (
	"fmt"
	"github.com/looplab/fsm"
)

func Ref{{.BizName}}StateMachine() *{{.BizName}}StateMachine {
	return stateMachine
}

type {{.BizName}}StateMachine struct {
    Fsm *fsm.FSM
}

var stateMachine = &{{.BizName}}StateMachine{}

func Init() {
	stateMachine.Fsm = fsm.NewFSM(
		"{{.InitState}}",
		fsm.Events{
		    {{ range .Events }}
			{Name: "{{.ActionName}}", Src: []string{ {{.SrcState}} }, Dst: "{{.DstState}}" },
			{{end}}
		},
		fsm.Callbacks{
			"leave_state":        func(event *fsm.Event) { fmt.Printf("now leaving state:%v to %v \n", event.Src, event.Dst) },
			 {{ range .State }}
			 "leave_{{.}}":       Exit{{.}}Func ,
			 "enter_{{.}}":       Enter{{.}}Func ,
			 {{end}}
		},
	)
}