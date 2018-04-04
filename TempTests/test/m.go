package main

import (
	"github.com/AndrienkoAleksandr/machine-exec/api/model"
	"sync"
	"fmt"
	"time"
)

type execs struct {
	mutex   *sync.Mutex
	execMap map[int]model.MachineExec
}

func main() {
	obj := execs{mutex: &sync.Mutex{}, execMap: make(map[int]model.MachineExec)}

	addElem(&obj)

	timer2 := time.NewTimer(time.Second)
	<-timer2.C

	obj.mutex.Lock()
	fmt.Println(obj.execMap)
	obj.mutex.Unlock()
}


func addElem(obj *execs) {
	go func() {
		for i:=10001; i < 200000; i++ {
			obj.mutex.Lock()
			obj.execMap[i] = model.MachineExec{ID:i}
			obj.mutex.Unlock()
		}
	}()

	for i:= 0; i < 100000; i++ {
		obj.mutex.Lock()
		obj.execMap[i] = model.MachineExec{ID:i}
		obj.mutex.Unlock()
	}
}
