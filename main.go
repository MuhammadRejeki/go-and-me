package main

import "fmt"

func main(){
	// Setiap const yang didefinisikan tidak harus di gunakan (berbeda dgn var)
	// Const tidak bisa deklarasi singkat menggunakan 
	
	const (
		firstName string = "Muhammad"
		lastName = "Rejeki"
		age = 27
	)

	fmt.Println(firstName,age)
}