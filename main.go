package main

import "fmt"
import "strings"
import "math"

func main() {
    var names = []string{"Muh", "Rejeki"}
    introducing("Halo, my name is", names)

	var diameter float64 = 15
    var area, circumference = lingkaran(diameter)
    fmt.Printf("luas lingkaran\t\t: %.2f \n", area)
    fmt.Printf("keliling lingkaran\t: %.2f \n", circumference)

	fSum := sumAll(10,15,30,44)
	fmt.Println("Sum :",fSum)

}

func introducing(message string, arr []string) {
    var nameString = strings.Join(arr, " ")
    fmt.Println(message, nameString)
}

func lingkaran(d float64) (float64, float64) {
    var area = math.Pi * math.Pow(d / 2, 2) // luas
    var circumference = math.Pi * d // keliling

    return area, circumference
}

func sumAll(argumens ...int) int{
	total := 0
	for _,v := range argumens{
		total += v
	}

	return total
}