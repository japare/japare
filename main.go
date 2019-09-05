package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {

	// parse config vars
	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "8000"
	}

	pathStatic, ok := os.LookupEnv("STATIC_FILES")
	if !ok {
		pathStatic = "dist/frontend/"
	}

	router := mux.NewRouter()

	// Create room for static files serving
	// mux.Handle("/node_modules/", http.StripPrefix("/node_modules", http.FileServer(http.Dir("./node_modules"))))
	// mux.Handle("/html/", http.StripPrefix("/html", http.FileServer(http.Dir("./html"))))
	// mux.Handle("/js/", http.StripPrefix("/js", http.FileServer(http.Dir("./js"))))
	// mux.Handle("/ts/", http.StripPrefix("/ts", http.FileServer(http.Dir("./ts"))))
	// mux.Handle("/css/", http.StripPrefix("/css", http.FileServer(http.Dir("./css"))))

	router.Handle("/{filename}", http.FileServer(http.Dir("./"+pathStatic)))

	// mux.HandleFunc("/api/login", api.Login)
	// mux.HandleFunc("/api/authenticate", api.Authenticate)

	// Any other request, we should render our SPA's only html file,
	// Allowing angular to do the routing on anything else other then the api
	// and the files it needs for itself to work.
	// Order here is critical. This html should contain the base tag like
	// <base href="/"> *href here should match the HandleFunc path below
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, pathStatic+"index.html")
	})

	log.Println("Serving at port ", port)
	log.Printf("Trying to find static files at: '%v'\n", pathStatic)
	log.Printf("Trying to find index html at: '%vindex.html'\n", pathStatic)

	log.Fatal(http.ListenAndServe("0.0.0.0:"+port, router))
}
