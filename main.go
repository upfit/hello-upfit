package main

import (
	"fmt"
	"net/http"

	cowsay "github.com/Code-Hex/Neo-cowsay"
)

func main() {
	say, err := cowsay.Say(
		cowsay.Phrase("Hello Upfit 💪"),
		cowsay.Type("default"),
		cowsay.BallonWidth(40),
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(say)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Hello Upfit 💪")
	})
	fmt.Println("Server running on port http://127.0.0.1:8080")
	http.ListenAndServe(":8080", nil)
}
