package human

import "fmt"

type gender int

const (
	Male gender = iota
	Female
)

// struct person
type Person struct {
	Name   string
	Age    int
	Gender gender
}

// to stirng method
func (p *Person) ToString() string {
	defer func() {
		if err := recover(); err != nil {
			
			fmt.Println(err)
		}
	}()
	
	var desc string
	if p.Gender == Female {
		desc = "Female"
	} else if Male == p.Gender {
		desc = "Male"
	} else {
		genderErr := GenderNotDefinedError{Message: "Gender not defined", Code: "100"}
		panic(genderErr)
	}
	return fmt.Sprintf("Name: %s, Age: %d, Gender: %s", p.Name, p.Age, desc)
}

// change name
func (p *Person) ChangeName(name string) {
	p.Name = name
}
