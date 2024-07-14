package main

import (
	"log"
	"net/http"
	"os"

	"github.com/7tsully-dev/exp-web/views"
)

func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}

func main() {

	views := []views.Page{
		views.ExpWebGraph{},
	}

	for _, v := range views {
		v.Render()
	}

	serverPages := "true"
	if len(os.Args) > 1 {
		serverPages = os.Args[1]
	}

	if serverPages == "false" {
		log.Println("Generated pages only, not server")
		return
	}
	fs := http.FileServer(http.Dir("examples/html"))
	log.Println("running server at http://localhost:8089")
	log.Fatal(http.ListenAndServe("localhost:8089", logRequest(fs)))
}
