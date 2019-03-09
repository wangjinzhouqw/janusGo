package main

import (
	"fmt"
	"runtime"
)

func main()  {
	funcName,fileName,line,ok := runtime.Caller(0)
	if ok{
		fmt.Println(fileName,":",line,",",runtime.FuncForPC(funcName).Name())
	}

}
