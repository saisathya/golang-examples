package struct_methods

import "fmt"

// Person ...
type Person struct {
	Name string
	Age  int
}

// setPersonPtr sets the name and age of a Person pointer
func (p *Person) setPersonPtr(name string, age int) {
	p.Name = name
	p.Age = age
}

// setPersonVal sets the name and age of a Person value
func (p Person) setPersonVal(name string, age int) {
	p.Name = name
	p.Age = age
}

func ExamplePtr() {
	personPtr := &Person{
		Name: "John",
		Age:  1,
	}
	fmt.Printf("personPtr is: %+v\n", personPtr)
	fmt.Printf("calling setPersonPtr on personPtr\n")
	personPtr.setPersonPtr("Steven Gerrard", 40)
	fmt.Printf("Value of personPtr after calling setPersonPtr is: %+v\n", personPtr)

	fmt.Printf("Calling setPersonVal on personPtr\n")
	personPtr.setPersonVal("Jamie Carragher", 41)
	fmt.Printf("Value of personPtr after calling setPersonVal is: %+v\n", personPtr)
}

func ExampleVal() {
	personVal := Person{
		Name: "John",
		Age:  1,
	}

	fmt.Printf("Value of personVal is: %+v\n", personVal)
	fmt.Printf("calling setPersonVal on personVal\n")
	personVal.setPersonVal("Steven Gerrard", 40)
	fmt.Printf("Value of personVal after calling setPersonVal is: %+v\n", personVal)

	fmt.Printf("Calling setPersonPtr on personVal\n")
	personVal.setPersonPtr("Jamie Carragher", 41)
	// This will work because golang will automatically convert the value to a pointer
	fmt.Printf("Value of personVal after calling setPersonVal is: %+v\n", personVal)
}
