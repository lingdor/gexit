package gexit

import "context"

type IGExit interface{
	GetContext() context.Context
	GetTask() IGExitTask
	GetNumber() int32
}
type IGExitTask interface {
	Done()
}

