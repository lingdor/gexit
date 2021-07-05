package gexit

import (
	"context"
	"log"
	"os"
	"sync/atomic"
)

type GExit struct{
	ctx context.Context
	number int32
	waitCh chan struct{}
	cancel context.CancelFunc
	signalCh chan os.Signal
}

func(ge *GExit)GetContext() context.Context{
	return ge.ctx
}

func (ge *GExit)GetTask() IGExitTask{
	atomic.AddInt32(&(ge.number),1)
	task:=new(GExitTask)
	task.parent=ge
	return task
}
func (ge *GExit)SubNumber(){
	atomic.AddInt32(&(ge.number),-1)
	if ge.GetNumber() == 0{
		select { //if empty, then wait done
			case ge.waitCh <- struct{}{}:break
			default:break
		}
	}
}


func (ge *GExit)GetNumber() int32 {
	return atomic.LoadInt32(&ge.number)
}
func (ge *GExit)SignalHandler() {
	defer func(){
		err:=recover()
		if err!= nil {
			log.Printf("%+v\n",err)
		}
	}()
	//if signal notified, close the context
	<- ge.signalCh
	ge.cancel() //context call
	ge.Wait() //block
	runPreExitFunc()  //preExit call
	//Exit complete
	os.Exit(0)
}

func (ge *GExit)Wait() {
	select {
		case <- ge.waitCh:return
		case <- ge.ctx.Done():return
	}
}

func NewIGExit(ctx context.Context)IGExit{
	cancelCtx,cancel := context.WithCancel(ctx)
	//init instance
	ins:=new (GExit)
	ins.signalCh = make(chan os.Signal,1)
	ins.cancel=cancel
	ins.ctx=cancelCtx
	ins.waitCh = make(chan struct{},0)
	//signal
	signalExitNotify(ins.signalCh)
	go ins.SignalHandler()
	//done
	return ins
}
