package main

import "fmt"

func main(){
	// Setiap const yang didefinisikan tidak harus di gunakan (berbeda dgn var)
	// Const tidak bisa deklarasi singkat menggunakan 
	
	const (
		firstName string = "Muhammad"
		lastName = "Rejeki"
		nick = "Kiikun"
		age = 27
	)

	fmt.Println(nick,age)


	// Type Declaration
	// membuat nama tipe sendiri

	type NIK string
	type stats bool 

	var myKTP NIK = "1111111"
	var married stats = true
	
	fmt.Println(myKTP,married)
	fmt.Println(NIK("2222222"),stats(false))
}