// gedung 1 3 lantai
// gedung 2 4 lantai
// gedung 3 1 lantai
// gedung 4 2 lantai
// gedung 5 7 lantai

package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	// gedung := []int{3, 4, 1, 2, 7}

	// for _, v := range gedung {
	// 	if v%2 == 0 {
	// 		fmt.Println("Ini genap", v)
	// 	} else {
	// 		fmt.Println("Ini ganjil", v)
	// 	}
	// }

	var (
		err error
		number int
	)
	v:="12"

	defer func(){
		log.Print(err)
	}()

	number, err = isNumber(v)
	if err != nil {
		return
	}
	fmt.Println(number)

}

func isNumber(v string)(num int, err error){
	num, err = strconv.Atoi(v)
	if err != nil{
		return
	}
	return
}
