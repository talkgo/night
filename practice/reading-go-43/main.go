package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

type Person struct{
	name string
}

func (p Person)Name()string  {
	return p.name
}
func changeName(p *Person, v string) {
	pointerVal := reflect.ValueOf(p)
	val := reflect.Indirect(pointerVal)
	member := val.FieldByName("name")
	ptrToName := unsafe.Pointer(member.UnsafeAddr())
	realPtrToName := (*string)(ptrToName)
	*realPtrToName = v
}
func main(){
	p := new(Person)
	changeName(p, "地狱咆哮.钮钴禄.刘本岩")
	fmt.Println(p.Name())
}
