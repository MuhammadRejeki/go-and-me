package main

import "fmt"

func main(){
	// Array
	days := [...]string{"senin","selasa","rabu","kamis","jumat","sabtu","minggu"}
	fmt.Println(days[2])

	// Slice 
	getDays		:= days[4:]
	getNew		:= []string{"lala","lulu"}
	getDays[0]	= "xjumat"
	fmt.Println(getDays)
	fmt.Println(getNew)
	fmt.Println(days)

	//Map
	book := map[string]int{}
	book["sejarah"] = 90
	book["math"] 	= 15

	data := map[string]string{
		"januari":"eropa",
		"februari":"asia", // koma diakhir ini wajib
	}
	fmt.Println(book); 
	fmt.Println(data); 
}