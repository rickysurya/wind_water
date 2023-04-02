package main

import "projek-pertama/chapter2/challenge1/routers"

func main() {
	var PORT = ":4000"

	routers.StartServer().Run(PORT)
}
