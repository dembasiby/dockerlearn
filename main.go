package main 

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	m := http.NewServeMux()

	m.HandleFunc("/", handlePage)

	const port = "8010"
	srv := http.Server{
		Handler: m,
		Addr: ":" + port,
		WriteTimeout: 30 * time.Second,
		ReadTimeout: 30 * time.Second,
	}

	// this blocks forever, until the Server
	// has an unrecoverable error
	err := srv.ListenAndServe()
	log.Fatal(err)
}

func handlePage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(200)
	
	const page = `<html>
	<head></head>
	<body>
		<p>Hello from Docker! I am a Go server.</p>
	</body>
	</html>`

	w.Write([]byte(page))
}
