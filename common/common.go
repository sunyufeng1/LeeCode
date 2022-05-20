package Common

import (
	"fmt"
	"reflect"
	"time"

	"leeCode.learn/testLee"
)

//动态调用方法
func InvokeObjectMethod(object interface{}, methodName string, args []reflect.Value) []reflect.Value {
	t1 := time.Now().UnixNano()
	result := reflect.ValueOf(object).MethodByName(methodName).Call(args)
	for i := 1; i < 10000; i++ {
		reflect.ValueOf(object).MethodByName(methodName).Call(args)
	}
	t2 := time.Now().UnixNano()
	runTime := (t2 - t1) / 1e6 //time.Since(t1)
	fmt.Println("run time : ", runTime, "毫秒")
	return result
}

//区别 这个不带统计
//args unite testBase 的参数
func InvokeObjectMethod1(object interface{}, methodName string) {
	inputs := make([]reflect.Value, 0)
	result := reflect.ValueOf(object).MethodByName(methodName).Call(inputs)
	leeObj := new(testLee.Obj)
	//finalResult testLee中方法返回的结果
	finalResult := InvokeObjectMethod(leeObj, methodName, result)
	fmt.Println("testBase result :")
	for _, value := range finalResult {
		fmt.Println(value)
	}
	//return result
}
