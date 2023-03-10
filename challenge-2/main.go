package main

import "fmt"

func main() {
	for i := 0; i <=4; i++ {
		fmt.Println("Nilai i =", i)
	}
	var characters = []rune{'\u0421', '\u0410', '\u0428', '\u0410', '\u0420', '\u0412', '\u041E'}
	for j :=0; j <=10; j++ {
		if j != 5{
			fmt.Println("Nilai j =", j)
		}
		if j == 5 {
			for index, value := range characters{
				fmt.Printf("character %U '%c' starts at byte position %d\n", value, value, index*2)
			}
		}
	}
}