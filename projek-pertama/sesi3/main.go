package main

import (
	"fmt"
)

type Person struct {
	name string
	age  int
}

func (p Person) Introduce(msg string) string {
	return fmt.Sprintf("%s my name is %s and I'm %d years old", msg, p.name, p.age)
}

func (p *Person) ChangeName2() {
	p.name = "Milo"
}

func main() {

	var person = &Person{name: "Airell", age: 25}
	person.ChangeName2()
	fmt.Println("Change name with ChangeName2 method", person.name)

	// var employee1 = struct {
	// 	person   Person
	// 	division string
	// }{
	// 	person:   Person{name: "Ananda", age: 23},
	// 	division: "Finance",
	// }

	// var numbers = []int{2,5,8,10,3,99,23}

	// var find = findOddNumers(numbers, func(number int) bool){

	// }

}

// func findOddNumbers(numbers []int, callback func(int) bool) int{
// 	var totalOddNumbers int
// 	for _, v:= range numbers {
// 		if callback(v){
// 			totalOddNumbers++
// 		}
// 	}
// 	return totalOddNumbers
// }

// func print(names ...string) []map[string]string {
// 	var result []map[string]string

// 	for i, v := range names {
// 		key := fmt.Sprintf("student%d", i+1)
// 		temp := map[string]string{
// 			key: v,
// 		}
// 		result = append(result, temp)

// 	}
// 	return result
// }

// func calculate(d float64) (float64, float64) {
// 	var area float64 = math.Pi * math.Pow(d/2, 2)

// 	var circumference = math.Pi * d

// 	return area, circumference
// }

// func greet(name, address string) {
// 	fmt.Printf("My name is %s, I live at %s", name, address)
// }

