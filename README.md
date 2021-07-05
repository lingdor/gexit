# gexit
graceful,safty exit your golang program.

# demo
/

```shell script
go get github.com/lingdor/gexit
```
init module code :
``` go
ctx:=context.Background()
gexit.InitGExit(ctx)
```
exit func callback
``` go
gexit.AddPreExitFunc(func(){
	fmt.Println("pre exit invoked!")
})
```

Context run
``` go
ctx:=gexit.GetContext()  //get exit context
for{
    select{
        case <-ctx.Done(): //when exit event
            fmt.Printf("goroutine %d at done!\n",num)
            return
        default:break
    }
    //normal program
}
```

# gexit-example

https://github.com/lingdor/gexit-example

#contact

mailto: bobby96333@gmail.com