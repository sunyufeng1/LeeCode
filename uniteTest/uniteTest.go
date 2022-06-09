package uniteTest

import (
	"bytes"
	"fmt"
	"reflect"
	"strconv"

	"github.com/sunyufeng1/LeeCode/testLee"
)

type Obj struct {
}

func (this *Obj) Test(id int, args ...interface{}) {

	var buf bytes.Buffer
	buf.WriteString("Test")
	buf.WriteString(strconv.Itoa(id))
	methodName := buf.String()

	callArs := make([]reflect.Value, len(args))
	for key, arg := range args {
		callArs[key] = reflect.ValueOf(arg)
	}
	leeObj := new(testLee.Obj)
	result := reflect.ValueOf(leeObj).MethodByName(methodName).Call(callArs)
	for _, value := range result {
		fmt.Println(value)
	}
}
