package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type Weather struct {
	Wind  int `json:"wind"`
	Water int `json:"water"`
}

func main() {
	for true {
		rand.Seed(time.Now().UnixNano())
		min := 1
		max := 100
		var data = map[string]interface{}{
			"water": rand.Intn((max - min + 1) + min),
			"wind":  rand.Intn((max - min + 1) + min),
		}
		requestJson, err := json.Marshal(data)
		client := &http.Client{}
		if err != nil {
			log.Fatalln(err)
		}

		req, err := http.NewRequest("POST", "https://jsonplaceholder.typicode.com/posts", bytes.NewBuffer(requestJson))
		req.Header.Set("Content-Type", "application/json")
		if err != nil {
			log.Fatalln(err)
		}

		res, err := client.Do(req)
		if err != nil {
			log.Fatalln(err)
		}

		//defer res.Body.Close()

		body, err := io.ReadAll(res.Body)

		if err != nil {
			log.Fatalln(err)
		}
		log.Println(string(body))

		var windWater Weather
		if err := json.Unmarshal([]byte(string(body)), &windWater); err != nil {
			fmt.Println(err)
			return
		}

		if windWater.Water < 5 {
			fmt.Println("Status Water : Aman")
		} else if windWater.Water >= 6 && windWater.Water <= 8 {
			fmt.Println("Status Water : Siaga")
		} else {
			fmt.Println("Status Water : Bahaya")
		}

		if windWater.Wind < 6 {
			fmt.Println("Status Wind  : Aman")
		} else if windWater.Wind >= 7 && windWater.Wind <= 15 {
			fmt.Println("Status Wind  : Siaga")
		} else {
			fmt.Println("Status Wind  : Bahaya")
		}
		time.Sleep(15 * time.Second)

	}
}
