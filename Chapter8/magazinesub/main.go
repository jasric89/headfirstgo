package main

import (
	"fmt"

	magazineinfo "github.com/jasric89/headfirstgo/Chapter8/magazinesub/magazineinfo"
)

func main() {
	var s magazineinfo.Subscriber
	employee := magazineinfo.Employee{Name: "Joy Carr", Salary: 6000}
	fmt.Println(employee.Name)
	fmt.Println(employee.Salary)
	address := magazineinfo.Address{Street: "123 Oak Street", City: "Omaha", State: "NE", PostalCode: "68111"}
	fmt.Println(address)
	s.Rate = 4.99
	fmt.Println(s.Rate)
	subscriber := magazineinfo.Subscriber{Name: "Aman Singh"}
	subscriber.Street = "123 Oak St"
	subscriber.City = "Omaha"
	subscriber.State = "NE"
	subscriber.PostalCode = "6811"
	fmt.Println("Subscriber Name:", subscriber.Name)
	fmt.Println("Street:", subscriber.Street)
	fmt.Println("City:", subscriber.City)
	fmt.Println("State:", subscriber.State)
	fmt.Println("Postal Code:", subscriber.PostalCode)

	employeetwo := magazineinfo.Employee{Name: "Jimminie Cricket"}
	employeetwo.Street = "456 Elm St"
	employeetwo.City = "Portland"
	employeetwo.State = "OR"
	employeetwo.PostalCode = "97222"
	fmt.Println("Employee Name:", employeetwo.Name)
	fmt.Println("Street:", employeetwo.Street)
	fmt.Println("City:", employeetwo.City)
	fmt.Println("Postal Code:", employeetwo.PostalCode)
}
