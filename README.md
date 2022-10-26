# gsrf
go struct reflection operation

## Introduce
It's a refection operation library, it focuses on providing simple 
and easy way to get reflection.

## Example
```go
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
	plist := GetStructPropertyListWithType(FinalExec{}, 	"CallBase")
	var finalExec = FinalExec{}
	finalExec.Fa = &FA{}
	finalExec.Fb = &FB{}
	var f NB
	ers := ExecMethod(&f, "Hi", "dd")
	if ers != nil {
		t.Error(ers)
	}
	for _, c := range plist {
		err := ExecMethod(GetInstanceFromFiledName(finalExec, c), 		"Hi")
		if err != nil {
			t.Error(err)
		}
	}
}
```

## License

​	MIT License

## Contribute

