package main

import (
	"projek-pertama/sesi7/routers"
)

func main() {
	var PORT = ":8080"

	routers.StartServer().Run(PORT)
}
