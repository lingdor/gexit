package gexit

import (
	"time"
)

type PreExitFunc func()
const LockTimeout=time.Second

var preExitFuncs []PreExitFunc = []PreExitFunc{}

func AddPreExitFunc(fn PreExitFunc){
	preExitFuncs = append(preExitFuncs,fn)
}

func runPreExitFunc(){
	for _,fn:=range preExitFuncs {
		fn()
	}
}