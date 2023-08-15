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
	s.Rate = 4.99
	fmt.Println(s.Rate)
}
