package gexit

import (
	"context"
	"sync"
	"testing"
)

const cnt = 10

func init(){
	InitGExit(context.Background())
}

func TestNumber(t *testing.T){

	wg:=&sync.WaitGroup{}
	wg.Add(cnt)
	tasks:=[cnt]IGExitTask{}

	for i:=0;i<cnt;i++ {
		go func(index int) {
			defer wg.Done()
			task:=GetTask()
			tasks[index]=task
		}(i)
	}
	wg.Wait()

	if int(GetNumber()) != cnt {
		t.Fail()
	}

	for i:=0;i<cnt;i++ {
		tasks[i].Done()
	}

	if int(GetNumber()) != 0 {
		t.Fail()
	}






}