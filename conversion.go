package main

import "fmt"

func main(){
	// how to do conversion

	var nilai32 int32 = 40000
	var nilai64 int64 = int64(nilai32)
	var nilai8 int8   = int8(nilai32) // jika melebihi kapasitas tipe integer maka perhitungan akan diawali dari awal lg

	fmt.Println(nilai32);
	fmt.Println(nilai64);
	fmt.Println(nilai8);


	var name = "Kiikun"
	var first byte = name[0] // dapet ascii
	var huruf = string(first) // dijadikan ke string dan conversion ke dari ascii

	fmt.Println(first,huruf)
}