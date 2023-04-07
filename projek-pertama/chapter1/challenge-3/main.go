package main

import (
	"fmt"
)

func main() {
	inpt := "selamat malam"
	charCount := make(map[string]int)

	for i := 0; i < len(inpt); i++ {
		fmt.Printf("%c\n", inpt[i])
		_, exist := charCount[fmt.Sprintf("%c", inpt[i])]
		if exist {
			charCount[fmt.Sprintf("%c", inpt[i])] += 1
		} else {
			charCount[fmt.Sprintf("%c", inpt[i])] = 1
		}

	}
	fmt.Println(charCount)

}
