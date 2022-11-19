package fsm

import (
	"fmt"
	"github.com/looplab/fsm"
)

{{ range . }}
func Enter{{.}}Func(event *fsm.Event) {
	fmt.Println("enter dst ",event.Dst,",modify your action logic here")
}
func Exit{{.}}Func(event *fsm.Event) {
	fmt.Println("Exit Src ",event.Src,",modify your action logic here")
}
{{end}}