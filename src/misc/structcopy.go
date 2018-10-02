package misc

import (
	"encoding/json"
	"fmt"
)

//
type Person struct {
	Id   int
	Name string
	Age  int
}

func main() {
	p := &Person{Id: 1, Name: "hell", Age: 24}
	b, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(b))
	var p1 Person
	json.Unmarshal(b, &p1)
	p1.Age = 56
	fmt.Println(p1)
	fmt.Printf("address of p: %p\n", &p)
	fmt.Printf("address of p1: %p\n", &p1)
	
	var p2 *Person
	p2 = &p1
	p2.Age = p.Age
	p2.Id = p.Id
	p2.Name = p.Name
	fmt.Printf("address of p2: %p\n", &p2)
	
}
