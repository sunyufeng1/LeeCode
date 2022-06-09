package Common

import (
	"fmt"
	"reflect"
	"time"

	"github.com/sunyufeng1/LeeCode/testLee"
)

//动态调用 testLee里的方法
func InvokeObjectMethod(methodName string, args []reflect.Value) []reflect.Value {
	leeObj := new(testLee.Obj)
	t1 := time.Now().UnixNano()
	result := reflect.ValueOf(leeObj).MethodByName(methodName).Call(args)
	for i := 1; i < 10000; i++ {
		reflect.ValueOf(leeObj).MethodByName(methodName).Call(args)
	}
	t2 := time.Now().UnixNano()
	runTime := (t2 - t1) / 1e6 //time.Since(t1)
	fmt.Println("run time : ", runTime, "毫秒")
	return result
}

//调用 uniteTest里的方法 并把结果 作为另一个地方的调用参数
func InvokeObjectMethod1(object interface{}, methodName string) {
	inputs := make([]reflect.Value, 0)
	result := reflect.ValueOf(object).MethodByName(methodName).Call(inputs)

	//finalResult testLee中方法返回的结果
	finalResult := InvokeObjectMethod(methodName, result)
	fmt.Println("testBase result :")
	for _, value := range finalResult {
		fmt.Println(value)
	}
	//return result
}
