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

	jim.updateName("jimmy")
	jim.print()
}

func (p person) updateName(newFirstName string){
	p.firstName = newFirstName
	// we expected jimmy to be printed but that didn't happen.
}

func (p person) print() {
	fmt.Printf("%+v", p)

}