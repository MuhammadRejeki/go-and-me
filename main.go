package main

import "fmt"

func main(){
	var point = 3840.0
	if percent := point / 100; percent >= 100 {
		fmt.Printf("%.1f%s perfect!\n", percent, "%")
	} else if percent >= 70 {
		fmt.Printf("%.1f%s good\n", percent, "%")
	} else {
		fmt.Printf("%.1f%s not bad\n", percent, "%")
	}

	var bilangan = 6
	switch bilangan {
		case 8:
			fmt.Println("perfect")
		case 7,6,5:
			fmt.Println("awesome")
		default:
			fmt.Println("not bad")
	}

}