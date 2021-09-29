package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

func main() {
	logrus.Info("Start server...")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		flusher, ok := w.(http.Flusher)
		if !ok {
			panic("expected http.ResponceWriter to be an http.Flusher")
		}
		w.Header().Set("X-Content-Type-Options", "nosniff")
		for i := 0; i < 10; i++ {
			fmt.Fprintf(w, "Chunk # %d \n", i)
			flusher.Flush()
			time.Sleep(500 * time.Millisecond)
		}
		fmt.Fprintf(w, "Bye")
		logrus.Info("Bye")
	})
	logrus.Info("Listening on localhost:8085")
	log.Fatal(http.ListenAndServe(":8085", nil))

}
