package main 

import "fmt"

func main(){
	var a = 10
	var b = 5

	
	fmt.Println("Math :",a + b, a - b, a * b)
	
	b++
	fmt.Println("Unary :",b)
	
	var stat = a == b // for integer u can use compare with using <= / >= 
	fmt.Println("Compare :",stat, !stat)

	var con1 bool = a > 50
	var con2 bool = a > 4
	var result bool = con1 || con2
	fmt.Println(result)
}