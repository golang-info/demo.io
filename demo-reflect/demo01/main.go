package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string "学生姓名"
	Age int `a:"1111"b:"3333"`
}

func main() {
	s := Student{}
	rt := reflect.TypeOf(s)
	fieldName, ok := rt.FieldByName("Name")

	if ok {
		fmt.Println(fieldName.Tag)
	}
	fieldAge, ok2 := rt.FieldByName("Age")

	if ok2 {
		fmt.Println(fieldAge.Tag.Get("a"))
		fmt.Println(fieldAge.Tag.Get("b"))
	}

	fmt.Println("type_Name: ", rt.Name())
	fmt.Println("type_NumField: ", rt.NumField())
	fmt.Println("Type_PkgPath: ", rt.PkgPath())
	fmt.Println("type_String: ", rt.String())

	fmt.Println("type.Kind.String: ", rt.Kind().String())
	fmt.Println("type.String()=", rt.String())

	//获取结构类型的字段名称
	for i := 0; i < rt.NumField(); i++ {
		fmt.Printf("type.Field[%d].Name:=%v \n", i, rt.Field(i).Name)
		fmt.Printf("type.Field[%d].Offset:=%v \n", i, rt.Field(i).Offset)
	}

	sc := make([]int, 10)
	sc = append(sc, 1, 2, 3)
	sct := reflect.TypeOf(sc)

	scet := sct.Elem()

	fmt.Println("slice element type.Kind()=", scet.Kind())
	fmt.Printf("slice element type.Kind()=%d\n", scet.Kind())
	fmt.Println("slice element type.String()=", scet.String())

	fmt.Println("slice element type.Name()=", scet.Name())
	fmt.Println("slice type.NumMethod()=", scet.NumMethod())
	fmt.Println("slice type.PkgPath()=", scet.PkgPath())
	fmt.Println("slice type.PkgPath()=", sct.PkgPath())
}
