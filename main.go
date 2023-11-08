package main

import "fmt"

func main(){
	var name string = "Kiikun" // atau bisa deklarasi singkat name := "Kiikun"
	var age int8	= 27 // jika tidak pakai type data akan auto deklari menggunakan int32 / string

	fmt.Println(name,age)

	var (
		firstName = "Muhammad"
		lastName = "Rejeki"
	)

	fmt.Println(firstName,lastName);

	// setiap var yang didefinisikan harus di gunakan atau golang akan menganggap itu error
}
