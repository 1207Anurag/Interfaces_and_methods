package main

import "fmt"

func main(){

}
type Contract struct{
	empId int
	Basic_Pay int
}
type Permanent struct{
	empId int
	Basic_Pay int
	pf int
}
type SalaryCalculator interface{
	calculateSalary() int
}
func(c Contract) calculateSalary()int{
return c.Basic_Pay
}
func(p Permanent)calculateSalary() int{
return p.Basic_Pay+p.pf
}
func calSalary(s []SalaryCalculator){
	sum:=0
	for _,v:=range s{
		sum=expense+v.calculateSalary)()
	}
}