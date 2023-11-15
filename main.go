package main

import "fmt"
import "strings"
import "math"

func main() {
    // example func with parameter
    var names = []string{"Muh", "Rejeki"}
    introducing("Halo, my name is", names)

    // example func return val multiple
	var diameter float64 = 15
    var area, circumference = lingkaran(diameter)
    fmt.Printf("luas lingkaran\t\t: %.2f \n", area)
    fmt.Printf("keliling lingkaran\t: %.2f \n", circumference)

    // example func variadic
	fSum := sumAll(10,15,30,44)
	fmt.Println("Sum :",fSum)

    // example func as parameters
    var data = []string{"wick", "jason", "ethan"}
    var dataContainsO = filter(data, func(each string) bool {
        return strings.Contains(each, "o")
    })
    var dataLenght5 = filter(data, func(each string) bool {
        return len(each) == 5
    })
    fmt.Println("data asli \t\t:", data)
    fmt.Println("filter ada huruf \"o\"\t:", dataContainsO)
    fmt.Println("filter jumlah huruf \"5\"\t:", dataLenght5)
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

func filter(data []string, callback func(string) bool) []string {
    var result []string
    for _, each := range data {
        if filtered := callback(each); filtered {
            result = append(result, each)
        }
    }
    return result
}

func main() {
    
}