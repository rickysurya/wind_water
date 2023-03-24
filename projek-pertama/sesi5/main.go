package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	go firstProcess(8)
	secondProcess(8)

	fmt.Println("No of goroutines:", runtime.NumGoroutine())

	time.Sleep(time.Second * 2)

	fmt.Println("main execution ended")
}

func goroutine() {
	fmt.Println("Hello")
}

func firstProcess(index int) {
	fmt.Println("First process func started")
	for i := 1; i <= index; i++{
		fmt.Println("i=", i)
	}
	fmt.Println("First process func ended")
}

func secondProcess(index int){
	fmt.Println("Second process func started")
	for j := 1; j <= index; j++{
		fmt.Println("j=", j)
	}
	fmt.Println("Second process func ended")
}

// ----------------------------------

// func main() {
// 	fruits := []string{"apple", "mango", "durian", "rambutan"}

// 	var wg sync.WaitGroup

// 	for index,fruit := range fruits {
// 		wg.Add(1)
// 		go printFruit(index, fruit, &wg)
// 	}
// 	wg.Wait()
// }

// func printFruit(index int, fruit string, wg *sync.WaitGroup){
// 	fmt.Printf("index => %d, fruit => %s\n", index, fruit)
// 	wg.Done()
// }

// --------------------

// func main(){
// 	c := make(chan string)

// 	go introduce("Airell", c)
// 	go introduce("Budi", c)
// 	go introduce("Caca", c)

// 	msg1 := <- c
// 	fmt.Println(msg1)

// 	msg2 := <- c
// 	fmt.Println(msg2)

// 	msg3 := <- c
// 	fmt.Println(msg3)

// 	close(c)
// }

// func introduce(student string, c chan string){
// 	result := fmt.Sprintf("Hai, my name is %s", student)

// 	c <-result
// }

// func main() {
// 	c := make(chan string)

// 	students := []string{"Airell", "Mailo", "Indah"}

// 	for _, v := range students {
// 		go func(student string) {
// 			fmt.Println("student", student)
// 			result := fmt.Sprintf("My name is %s", student)
// 			c <- result
// 		}(v)
// 	}

// 	for i := 1; i <= 3; i++ {
// 		print(c)
// 	}

// 	close(c)
// }

// func print(c chan string) {
// 	fmt.Println(<-c)
// }


