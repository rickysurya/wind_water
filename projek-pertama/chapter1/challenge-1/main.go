package main

import (
	"fmt"
)

func main() {

	// menampilakan nilai i : 21
	i := 21
	fmt.Println(i)

	// menampilkan tipe data dari variabel i
	fmt.Printf("%T\n", i)

	// menampilkan tanda %
	fmt.Println("%")

	// menampilkan nilai boolean j : true
	j := true
	fmt.Println(j)

	// menampilkan nilai boolean j : false
	fmt.Println(j && false)

	// menampilkan unicode russia : Я (ya)
	fmt.Printf("%q\n", "Я")

	// menampilkan nilai base 10 : 21
	fmt.Printf("%d\n", 21)

	// menampilkan nilai base 8 :25
	fmt.Printf("%o\n", 25)

	// menampilkan nilai base 16 : f
	fmt.Printf("%x\n", "f")

	// menampilkan nilai base 16 : F
	fmt.Printf("%X\n", "F")

	// menampilkan unicode karakter Я : U+042F
	fmt.Printf("%+q\n", "Я")

	// menampilkan float : 123.456000
	var k float64 = 123.456
	fmt.Printf("%.6f\n", k)

	// menampilkan float scientific : 1.234560E+02
	fmt.Printf("%E", k)

}
