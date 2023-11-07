package main

import "fmt"

func main(){
	fmt.Println("Tipe Number : int (-+), uint(+), float[32&64], bool(true/false)");
	fmt.Println("Alias yang biasa digunakan pada number yaitu byte (uint8), rune(int32), int(int32), uint(uint32) \n");

	fmt.Println("Tipe Text hanya menggunakan kode String dan tanda kutip 2 (''). contoh:\n");
	fmt.Println(len("Kiikun")); // fungsi total digit string
	fmt.Println("Muhammad Rejeki"[0]); // mengambil kode asci di huruf pertama
	fmt.Println("Muhammad Rejeki Kun"[1]);
}