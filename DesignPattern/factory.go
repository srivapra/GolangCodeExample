package main

import "fmt"

type Emplyee struct {
	EmpId   string
	EmpName string
}

func factory(empId string, empName string) Emplyee {
	return Emplyee{
		EmpId:   empId,
		EmpName: empName,
	}

}

func main() {
	fmt.Print(factory("123", "Prashant"))
}
