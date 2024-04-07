/*
Now let’s look at a practical use of interface.

We will write a simple program that calculates the total expense for a company based on the
individual salaries of the employees. For brevity, we have assumed that all expenses are in USD
*/

package main

import "fmt"

type SalaryCalculator interface {
	SalaryCalculate() int
}

type Permanent struct {
	empId    int
	basicPay int
	pf       int
}

type Contract struct {
	empId    int
	basicPay int
}

type Freelancer struct {
	empId       int
	ratePerHour int
	totalHours  int
}

func (p Permanent) SalaryCalculate() int {
	return p.basicPay + p.pf
}

func (c Contract) SalaryCalculate() int {
	return c.basicPay
}

func (f Freelancer) SalaryCalculate() int {
	return f.ratePerHour * f.totalHours
}

func totalExpense(s []SalaryCalculator) {
	expense := 0
	for _, v := range s {
		expense = expense + v.SalaryCalculate()

	}
	fmt.Printf("Total expense are %d ", expense)
}

func main() {
	pemp1 := Permanent{
		empId:    1,
		basicPay: 5000,
		pf:       20,
	}
	pemp2 := Permanent{
		empId:    2,
		basicPay: 6000,
		pf:       30,
	}
	cemp1 := Contract{
		empId:    3,
		basicPay: 3000,
	}

	// freelancer1 := Freelancer{
	// 	empId:       4,
	// 	ratePerHour: 70,
	// 	totalHours:  120,
	// }
	// freelancer2 := Freelancer{
	// 	empId:       5,
	// 	ratePerHour: 100,
	// 	totalHours:  100,
	// }

	// We can pass freelancer1, freelancer2 in SalaryCalculator

	s := []SalaryCalculator{pemp1, pemp2, cemp1}
	totalExpense(s)
}

/*
We have two kinds of employees in the company, Permanent and Contract defined by structs
The salary of permanent employees is the sum of the basicpay and pf whereas for contract employees
it’s just the basic pay basicpay. This is expressed in the corresponding CalculateSalary methods
By declaring this method, both Permanent and Contract structs now implement the SalaryCalculator
interface.

The totalExpense function declared expresses the beauty of interfaces. This method takes a slice
of SalaryCalculator interface []SalaryCalculator as a parameter. we pass a slice that
contains both Permanent and Contract types to the totalExpense function. The totalExpense function
calculates the expense by calling the CalculateSalary method of the corresponding type.

*/

/*
The biggest advantage of this is that totalExpense can be extended to any new employee type
without any code changes. Let’s say the company adds a new type of employee Freelancer
with a different salary structure. This Freelancer can just be passed in the slice argument
to totalExpense without even a single line of code change to the totalExpense function.This method
will do what it’s supposed to do as Freelancer will also implement the SalaryCalculator interface

*/
