package fsm

import (
	"fmt"
	"github.com/looplab/fsm"
)

func EntersightseerFunc(event *fsm.Event) {
	fmt.Println("enter dst ", event.Dst, ",modify your action logic here")
}
func ExitsightseerFunc(event *fsm.Event) {
	fmt.Println("Exit Src ", event.Src, ",modify your action logic here")
}

func EnterbuyerFunc(event *fsm.Event) {
	fmt.Println("enter dst ", event.Dst, ",modify your action logic here")
}
func ExitbuyerFunc(event *fsm.Event) {
	fmt.Println("Exit Src ", event.Src, ",modify your action logic here")
}

func EnterholderFunc(event *fsm.Event) {
	fmt.Println("enter dst ", event.Dst, ",modify your action logic here")
}
func ExitholderFunc(event *fsm.Event) {
	fmt.Println("Exit Src ", event.Src, ",modify your action logic here")
}
