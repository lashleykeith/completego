package main


import "fmt"

type contactInfo struct {
	email string
	zipCode int
}

type person struct {
	firstName string
	lastName string
	contactInfo //contactInfo has a field name contactinfo 
}

func main(){
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email: "jim@gmail.com",
			zipCode: 94000,
		},
	}

	jimPointer := &jim
	jimPointer.updateName("jimmy")
	jim.print()
}

func (pointerToPerson *person) updateName(newFirstName string){
	(*pointerToPerson).firstName = newFirstName
	// *person This is type description it means we're working with a pointer to a person.
	// *pointerToPerson This is an operator it means we want to manipulate the value the pointer is referencing.
	// Turn addres into value with *address
	// Turn value into address with &value
}
func (p person) print() {
	fmt.Printf("%+v", p)

}