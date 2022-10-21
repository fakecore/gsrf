package gsrf

import (
	"fmt"
	"testing"
)

type TestF struct {
}

func (b TestF) HelloSimple() {
	fmt.Println("hiv")
}

func (b *TestF) HelloPointer(s string) {
	fmt.Println("hi,this is hi", s)
}

func (b *TestF) helloLower(s string) {
	fmt.Println("hi,", s)
}

func TestOperation(t *testing.T) {
	//StructFunctionList(&TestF{})
	err := ExecMethod(&TestF{}, "HelloPointer", "eeeee", "ddddd")
	if err != nil {
		fmt.Println(err.Error())
	}
	err = ExecMethod(&TestF{}, "HelloSimple", "eeeee", "ddddd")
	if err != nil {
		fmt.Println(err.Error())
	}
	err = ExecMethod(&TestF{}, "NooooV", "eeeee", "ddddd")
	if err != nil {
		fmt.Println(err.Error())
	}
	err = ExecMethod(&TestF{}, "helloLower", "eeeee")
	if err != nil {
		fmt.Println(err.Error())
	}
}

type SrcData struct {
	v1   int
	Test string
}
type DestData struct {
	Uk   int
	v1   int
	Test string
}

func TestStructCopy(t *testing.T) {
	s := SrcData{v1: 23, Test: "hhhhh"}
	var d DestData
	StructCopy(&s, &d)
	fmt.Printf("%v %v", s, d)
	if s.v1 == d.v1 || s.Test != d.Test {
		t.Error("copy failed")
	}
}
