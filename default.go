package gexit

import (
	"context"
)
var Default IGExit = new(GExitNil)

func GetNumber() int32{
	return Default.GetNumber()
}

func GetContext() context.Context{
	return Default.GetContext()
}
func GetTask() IGExitTask{
	return Default.GetTask()
}

func InitGExit(ctx context.Context){
	Default=NewIGExit(ctx)
}