package gexit

import "sync/atomic"

type GExitTask struct {
	parent *GExit
	subed int32
}

func (get *GExitTask)Done(){
	if atomic.CompareAndSwapInt32(&(get.subed),0,1){
		get.parent.SubNumber()
	}
}

