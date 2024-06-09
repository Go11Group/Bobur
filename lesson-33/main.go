package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func main() {

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/ismFamilya", name_lastname)
	http.HandleFunc("/karrajadval", karraJadval)
	http.HandleFunc("/time", timeHandler)
	http.HandleFunc("/random", newRandom)


	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error Server...", err)
		return
	}

}

func hello(w http.ResponseWriter, r *http.Request) {

	message := "Hello world!\n"

	w.Write([]byte(message))
}

func name_lastname(w http.ResponseWriter, r *http.Request) {
	name := "Bobur"
	lastname := "Yusupov"

	w.Write([]byte(name + " " + lastname))
}

func karraJadval(w http.ResponseWriter, r *http.Request) {
	for i := 1; i <= 10; i++ {
		for j := 1; j <= 10; j++ {
			a := i * j
			chiqish := fmt.Sprintf("%d * %d = %d\n	", i, j, a)
			w.Write([]byte(chiqish))
		}
		w.Write([]byte{'\n'})
	}
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	cur := time.Now()
	format := cur.Format("2006-01-02 15:04:05")

	w.Write([]byte(format))
}

func newRandom(w http.ResponseWriter, r *http.Request) {
	rnd := rand.Intn(100)
	fmt.Fprintf(w, "%d", rnd)
}
