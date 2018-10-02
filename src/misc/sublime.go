package misc

import (
	"fmt"
	
	"human"
)

func main() {
	
	person := human.Person{Name: "ethan.li", Age: 30, Gender: human.Male}
	
	fmt.Println(person.ToString())
	
	person.ChangeName("feng")
	
	fmt.Println(person.ToString())
}
