package main

import (
	"encoding/json"
	"fmt"
)

// type Employee struct {
// 	FullName string `json:"full_name"`
// 	Email    string `json:"email"`
// 	Age      int    `json:"age"`
// }

// func main() {
// 	var jsonString = `[
// 			{
// 				"full_name": "Airell",
// 				"email": "airell@mail.com",
// 				"age": 23
// 			},
// 			{
// 				"full_name": "ye",
// 				"email": "ye@mail.com",
// 				"age": 40
// 			}
// 		]	
// 	`

// 	// var result Employee

// 	// var temp interface{}

// 	var result []Employee

// 	var err = json.Unmarshal([]byte(jsonString), &result)

// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return
// 	}
// 	// var result = temp.(map[string]interface{})

// 	for i, v := range result{
// 		fmt.Printf("Index %d: %+v\n", i+1, v)
// 	}

// 	// fmt.Println("full_name", result["full_name"])
// 	// fmt.Println("email", result["email"])
// 	// fmt.Println("age", result["age"])

// }


type User struct {
	FullName	string `json:"Name"`
	Age			int
}

func main(){
	var object = []User{{"john wick", 27}, {"ethan", 32}}

	var jsonData, err = json.Marshal(object)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var jsonString = string (jsonData)
	fmt.Println(jsonString)
}