package uniteTest

import (
	"testing"
)

func Test0(t *testing.T) {
	leeObj := new(Obj)
	leeObj.Test(0, "testtt")
}

func Test1(t *testing.T) {
	testObj := new(Obj)
	nums := []int{2, 7, 11, 15}
	target := 26
	testObj.Test(1, nums, target)
}

func Test2(t *testing.T) {
	testObj := new(Obj)
	testObj.Test(2, 12345678)
}

func Test3(t *testing.T) {
	testObj := new(Obj)
	testObj.Test(3, 10)
}

func Test4(t *testing.T) {
	testObj := new(Obj)
	testObj.Test(4, "MCMXCIV")
}

func Test5(t *testing.T) {
	testObj := new(Obj)
	testObj.Test(5, []string{"a2", "a3"})
}

func Test6(t *testing.T) {
	testObj := new(Obj)
	testObj.Test(6, "()[]{}")
}
