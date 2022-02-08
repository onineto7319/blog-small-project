package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("test"))
	})
	http.ListenAndServe(":" + os.Getenv("PORT"), nil)
}
