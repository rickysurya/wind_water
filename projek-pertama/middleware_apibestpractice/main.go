package main

import(
	"net/http"
	"log"
	"fmt"
)



func main(){
	mux := http.NewServeMux()

	endpoint := http.HandlerFunc(greet)

	mux.Handle("/", middleware1(middleware2(endpoint)))

	fmt.Println("Listening to port 8080")

	err := http.ListenAndServe(":8080", mux)

	log.Fatal(err)
}

func greet(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Hello World!"))
}

func middleware1(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		log.Println("middleware pertama")
		next.ServeHTTP(w, r)
	})
}
func middleware2(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		log.Println("middleware kedua")
		next.ServeHTTP(w, r)
	})
}

