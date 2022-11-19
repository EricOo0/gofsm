package main

import (
	"encoding/json"
	"fmt"
	"gofsm/fsm"
	"io/ioutil"
	"os"
	"text/template"
)

type StateMachineCfg struct {
	Name      string
	InitState string `json:"init_state"`
	Event     []StateEvent
}
type StateEvent struct {
	Name string
	Src  []string
	Dst  string
}

type StateMachineTemplate struct {
	PackageName string
	BizName     string
	InitState   string
	Events      []MachineEvent
	State       []string
}

type MachineEvent struct {
	ActionName string
	SrcState   string
	DstState   string
}

type CallBackFunc struct{}

func buildFsm() {
	// 打开配置文件
	jsonFile, err := os.Open("config.json")
	if err != nil {
		panic(err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	// 读取配置
	fsmCfg := &StateMachineCfg{}
	err = json.Unmarshal(byteValue, fsmCfg)
	if err != nil {
		panic(err)
	}

	// 读取模版文件
	tpl, err := template.ParseFiles("./template/state_machine.tpl")
	if err != nil {
		panic(err)
	}
	actionTpl, err := template.ParseFiles("./template/action.tpl")
	if err != nil {
		panic(err)
	}

	eventTpl := &StateMachineTemplate{
		PackageName: "fsm",
		BizName:     "Stock",
		InitState:   fsmCfg.InitState,
	}
	stateMap := make(map[string]bool)
	for _, event := range fsmCfg.Event {
		src, err := json.Marshal(event.Src)
		if err != nil {
			fmt.Println(fmt.Sprintf("add event error,src:%v,dst:%v,err:%v", event.Src, event.Dst, err))
			continue
		}

		for _, src := range event.Src {
			if _, ok := stateMap[src]; !ok {
				eventTpl.State = append(eventTpl.State, src)
				stateMap[src] = true
			}
		}
		if _, ok := stateMap[event.Dst]; !ok {
			eventTpl.State = append(eventTpl.State, event.Dst)
			stateMap[event.Dst] = true
		}
		eventTpl.Events = append(eventTpl.Events, MachineEvent{
			ActionName: event.Name,
			SrcState:   string(src)[1 : len(string(src))-1],
			DstState:   event.Dst,
		})
	}
	//生产状态机
	file, err := os.Create("./fsm/state_machine.go")
	if err != nil {
		panic(err)
	}
	tpl.Execute(file, eventTpl)

	file2, err := os.Create("./fsm/action.go")
	if err != nil {
		panic(err)
	}
	actionTpl.Execute(file2, eventTpl.State)

}
func main() {
	buildFsm()

	fsm.Init()
	err := fsm.RefStockStateMachine().Fsm.Event("openAccount")
	fmt.Println(err)
}
