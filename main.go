package main

import "fmt"

func main(){

	
	var xs = "7878" // string
	for i, v := range xs {
		fmt.Println("Index=", i, "Value=", v)
		// v get data ascii dari tiap digit string nya
	}

	var ys = [5]int{10, 20, 30, 40, 50} // array
	for _, v := range ys {
		fmt.Println("Value=", v)
	}

	for i := 1; i <= 10; i++ {
		if i % 2 == 1 {
			continue //skip tahap sekarang dan lanjut keberikutnya
		}
	
		if i > 8 {
			break
		}
	
		fmt.Println("Angka", i)
	}

	loopawal: //menamakan loop
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == 3 {
				break loopawal
			}
			fmt.Print("matriks [", i, "][", j, "]", "\n")
		}
	}
	
}