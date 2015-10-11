package function

import (
	"testing"
)

func test1(t *testing.T, hh string) {
	t.Log("i'm here")
	t.Log(hh)
}

func TestCall(t *testing.T) {
	f := Bind(test1, PH(P_1), "haha")
	f.Call(t)
}

func test2(t *testing.T, hh string) {
	t.Log("i'm here")
	t.Log(hh)
}

func TestCall2(t *testing.T) {
	f := Bind(test1, PH(P_1), PH(P_2))
	f.Call(t, "test2")
}
