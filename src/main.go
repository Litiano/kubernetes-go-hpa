package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
)

func greeting(s string) string {
	return fmt.Sprintf("<b>%s</b>", s)
}

func handler(w http.ResponseWriter, r *http.Request) {
	x := 0.0001
	for i := 0; i <= 10000000000; i++ {
		x += math.Sqrt(x)
	}
	fmt.Fprint(w, greeting("Code.education Rocks!"))
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":80", nil))
}