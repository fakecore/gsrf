package gsrf

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"reflect"
	"sort"
	"testing"
)

type TestF struct {
	VInt    uint
	VString string
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
	err := ExecMethod(&TestF{}, "HelloPointer", "eeeee")
	assert.Equal(t, err == nil, true)
	if err != nil {
		t.Error(err)
	}
	err = ExecMethod(&TestF{}, "HelloSimple", "eeeee", "ddddd")
	assert.Equal(t, err != nil, true)
	err = ExecMethod(&TestF{}, "NooooV", "eeeee", "ddddd")
	assert.Equal(t, err != nil, true)
	err = ExecMethod(&TestF{}, "helloLower", "eeeee")
	//becuase of the `helloLower` is not exported
	assert.Equal(t, err != nil, true)
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
func TestGetStructFiled(t *testing.T) {
	tList := GetStructPropertyList(TestF{})
	targetPList := []string{"VInt", "VString"}
	sort.Strings(tList)
	sort.Strings(targetPList)
	assert.Equal(t, targetPList, tList)
}
func TestGetStructFiledWithType(t *testing.T) {
	tList := GetStructPropertyListWithType(TestF{}, reflect.TypeOf("").Name())
	targetPList := []string{"VString"}
	sort.Strings(tList)
	sort.Strings(targetPList)
	assert.Equal(t, targetPList, tList)
}

type CallBase interface {
	Hi()
}

type FA struct {
}

func (f *FA) Hi() {
	fmt.Println("hi from FA")
}

type FB struct {
}

func (f *FB) Hi() {
	fmt.Println("hi from FB")
}

type NB struct {
}

func (f *NB) Hi(s string, e string, ddd int, sd int) {
	fmt.Println("hi from NB")
}

type FinalExec struct {
	Fa CallBase
	Fb CallBase
}

func TestExecMemberVariableMethod(t *testing.T) {
	plist := GetStructPropertyListWithType(FinalExec{}, "CallBase")
	var finalExec = FinalExec{}
	finalExec.Fa = &FA{}
	finalExec.Fb = &FB{}
	var f NB
	ers := ExecMethod(&f, "Hi", "dd")
	if ers != nil {
		t.Error(ers)
	}
	for _, c := range plist {
		err := ExecMethod(GetInstanceFromFiledName(finalExec, c), "Hi")
		if err != nil {
			t.Error(err)
		}
		//GetInstanceFromFiledName(finalExec, c)
	}
}
