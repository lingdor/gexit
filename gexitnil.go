package gexit

import (
	"context"
)

type GExitNil struct{
}

func(ge *GExitNil)GetContext() context.Context{
	panic("GExit is not initialized")
}

func (ge *GExitNil)GetTask() IGExitTask{
	panic("GExit is not initialized")
}
func (ge *GExitNil)SubNumber(){
	panic("GExit is not initialized")
}


func (ge *GExitNil)GetNumber() int32 {
	panic("GExit is not initialized")
}

func (ge *GExitNil)Wait() {
	panic("GExit is not initialized")
}
