package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Hello World")
		d, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(rw, "Ooops", http.StatusBadRequest)
			// rw.WriteHeader(http.StatusBadRequest);
			// rw.Write([]byte("Ooops"))
			return
		}

		fmt.Fprintf(rw, "Hello %s", d)
	})

	http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request) {
		log.Println("GoodBye!")
	})

	http.ListenAndServe(":9090", nil)
}
