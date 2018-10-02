package misc

import (
	"fmt"
	"reflect"
)

//
type User struct {
	Id   int64
	Name string
	Age  int
}

func (this *User) ToString() {
	fmt.Printf("id:%d, name: %s, Age:%d\n", this.Id, this.Name, this.Age)
}

//
func (this User) SayHello() {
	fmt.Printf("Hello, my name is %s\n", this.Name)
}

//
func (this *User) SetAge(age int) {
	this.Age = age
	
}

func main() {
	u := User{Id: 1, Name: "Ethan", Age: 11}
	v := reflect.ValueOf(&u)
	method := v.MethodByName("ToString")
	// args := []reflect.Value{reflect.ValueOf(12)}
	method.Call(nil)
	fmt.Println(u)
}
