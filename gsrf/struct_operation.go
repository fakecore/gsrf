package gsrf

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

func GetStructName(src interface{}) (name string) {
	return reflect.TypeOf(src).Elem().Name()
}

// StructFunctionList must pass the struct pointer to get all the methods of src struct
func GetStructFunctionList(src interface{}) (fnList []string) {
	typeList := reflect.TypeOf(src)
	count := typeList.NumMethod()
	for i := 0; i < count; i++ {
		fnList = append(fnList, typeList.Method(i).Name)
	}
	return fnList
}

func GetStructPropertyList(src interface{}) (fnList []string) {
	typeList := reflect.TypeOf(src)
	count := typeList.NumField()
	for i := 0; i < count; i++ {
		fnList = append(fnList, typeList.Field(i).Name)
	}
	return fnList
}

func GetStructPropertyListWithType(src interface{}, propertyType string) (fieldList []string) {
	typeList := reflect.TypeOf(src)
	count := typeList.NumField()
	for i := 0; i < count; i++ {
		if getPureFiledName(typeList.Field(i).Type.String()) == propertyType {
			fieldList = append(fieldList, typeList.Field(i).Name)
		}
	}
	return fieldList
}
func getPureFiledName(name string) string {
	nlist := strings.Split(name, ".")
	return nlist[len(nlist)-1]
}

//GetInstanceFromFiledName
//@param src: struct pointer
//@param fieldName:member variable name
func GetInstanceFromFiledName(src interface{}, fieldName string) any {
	s := reflect.ValueOf(src).FieldByName(fieldName).Elem().Interface()
	return s
}

//ExecMethod
//@param: src target struct.Should pass pointer type
//@param: fnName call function name
//@param: args function execute with those params
//only supports exported function
//fixme: check the count and type of args are matching target function
func ExecMethod(src interface{}, fnName string, args ...interface{}) (err error) {
	typeList := reflect.TypeOf(src)
	fn, isOK := typeList.MethodByName(fnName)
	if !isOK {
		err = errors.New(fmt.Sprintf("not found:%s in %s\n", fnName, typeList.Elem().Name()))
		return
	}
	var params []reflect.Value
	for _, c := range args {
		params = append(params, reflect.ValueOf(c))
	}
	fnCount := fn.Type.NumIn() - 1
	if fnCount != len(params) {
		err = errors.New(fmt.Sprintf("param count error,target:%d,now:%d\n", fnCount, len(params)))
		return
	}
	reflect.ValueOf(src).MethodByName(fnName).Call(params)
	return nil
}

// StructCopy all action is shallow copy
// from https://juejin.cn/post/6844904009505964039
func StructCopy(src interface{}, dst interface{}) (err error) {
	// 防止意外panic
	defer func() {
		if e := recover(); e != nil {
			err = errors.New(fmt.Sprintf("%v", e))
		}
	}()

	dstType, dstValue := reflect.TypeOf(dst), reflect.ValueOf(dst)
	srcType, srcValue := reflect.TypeOf(src), reflect.ValueOf(src)

	// dst必须结构体指针类型
	if dstType.Kind() != reflect.Ptr || dstType.Elem().Kind() != reflect.Struct {
		return errors.New("dst type should be a struct pointer")
	}

	// src必须为结构体或者结构体指针
	if srcType.Kind() == reflect.Ptr {
		srcType, srcValue = srcType.Elem(), srcValue.Elem()
	}
	if srcType.Kind() != reflect.Struct {
		return errors.New("src type should be a struct or a struct pointer")
	}

	// 取具体内容
	dstType, dstValue = dstType.Elem(), dstValue.Elem()

	// 属性个数
	propertyNums := dstType.NumField()

	for i := 0; i < propertyNums; i++ {
		// 属性
		property := dstType.Field(i)
		// 待填充属性值
		propertyValue := srcValue.FieldByName(property.Name)

		// 无效，说明src没有这个属性 || 属性同名但类型不同
		if !propertyValue.IsValid() || property.Type != propertyValue.Type() {
			continue
		}

		if dstValue.Field(i).CanSet() {
			dstValue.Field(i).Set(propertyValue)
		}
	}

	return nil
}
