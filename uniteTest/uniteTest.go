package uniteTest

import (
	"bytes"
	"runtime"
	"strconv"
	"strings"

	Common "github.com/sunyufeng1/LeeCode/common"
	"github.com/sunyufeng1/LeeCode/testLee"
)

type Obj struct {
	funObj *testLee.Obj
}

func (this *Obj) Test(id int) {
	this.funObj = new(testLee.Obj)
	_ = this.funObj
	var buf bytes.Buffer
	buf.WriteString("Test")
	buf.WriteString(strconv.Itoa(id))
	str := buf.String()
	Common.InvokeObjectMethod1(this, str) //第二个参数是测试用例的方法名
}

func (this *Obj) getFunName() string {
	pc, file, line, ok := runtime.Caller(1)
	_ = file
	_ = line
	_ = ok
	f := runtime.FuncForPC(pc)
	names := strings.Split(f.Name(), ".")
	funcName := names[len(names)-1]
	return funcName
}

//入口测试
func (this *Obj) Test0() string {
	return "test0"
}

func (this *Obj) Test1() ([]int, int) { //[]interface{}{
	//funcName := this.getFunName()
	nums := []int{2, 7, 11, 15}
	target := 26
	//fmt.Println(target,nums)
	//return arr
	return nums, target
}

func (this *Obj) Test2() int {
	return 12345678
}

func (this *Obj) Test3() int {
	return 10
}

func (this *Obj) Test4() string {
	return "MCMXCIV"
}

func (this *Obj) Test5() []string {
	return []string{"a2", "a3"}
}

func (this *Obj) Test6() string {
	return "()[]{}"
}
