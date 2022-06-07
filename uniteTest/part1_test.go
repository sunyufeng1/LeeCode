package uniteTest

import (
	"testing"

	"github.com/sunyufeng1/LeeCode/testLee"
)

func Test0(t *testing.T) {
	testObj := new(Obj)
	testObj.Test(0)
}

func Test00(t *testing.T) {
	leeObj := new(testLee.Obj)
	leeObj.Test0("testtt")
}

func Test1(t *testing.T) {
	testObj := new(Obj)
	testObj.Test(1)
}

func Test2(t *testing.T) {
	testObj := new(Obj)
	testObj.Test(2)
}
