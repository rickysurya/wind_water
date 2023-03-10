package main

import "fmt"

func main() {
	// var curr = 2021
	// if age:= curr - 1998; age < 17{
	// 	fmt.Println("Dah bole buat sim")
	// } else {
	// 	fmt.Println("Belum bole buat sim")
	// }

	// var score = 8
	// switch score {
	// case 8:
	// 	fmt.Println("perfect")
	// case 7:
	// 	fmt.Println("awesome")
	// }

	// //FizzBuzz
	// input := 5
	// //cetak "FIZZ" kalau habis dibagi 3
	// // kalau habis dibagi 5 cetak "BUZZ"
	// // kalau habis dibagi 3 dan 5 cetak "FIZZBUZZ"

	// if (input%5 == 0) && (input%3 == 0) {
	// 	fmt.Println("FIZZBUZZ")
	// } else if input%5 == 0 {
	// 	fmt.Println("BUZZ")
	// } else if input%3 == 0 {
	// 	fmt.Println("FIZZ")
	// }

	// if (input % 3 == 0){
	// 	fmt.Print("FIZZ")
	// }

	var numbers [4]int
	numbers = [4]int{1,2,3,4}
	var strings = [3]string{"Airell", "Milo", "Nanda"}

	fmt.Printf("%#v\n", numbers)
	fmt.Printf("%#v\n", strings)

}
