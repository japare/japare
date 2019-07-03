package main

import (
	"net/http"
	"log"
)

func main() {
	// Creates a new serve mux
	mux := http.NewServeMux()

	// Create room for static files serving
	// mux.Handle("/node_modules/", http.StripPrefix("/node_modules", http.FileServer(http.Dir("./node_modules"))))
	// mux.Handle("/html/", http.StripPrefix("/html", http.FileServer(http.Dir("./html"))))
	// mux.Handle("/js/", http.StripPrefix("/js", http.FileServer(http.Dir("./js"))))
	// mux.Handle("/ts/", http.StripPrefix("/ts", http.FileServer(http.Dir("./ts"))))
	// mux.Handle("/css/", http.StripPrefix("/css", http.FileServer(http.Dir("./css"))))

	
	// mux.HandleFunc("/api/login", api.Login)
	// mux.HandleFunc("/api/authenticate", api.Authenticate)

	// Any other request, we should render our SPA's only html file,
	// Allowing angular to do the routing on anything else other then the api    
	// and the files it needs for itself to work.
	// Order here is critical. This html should contain the base tag like
	// <base href="/"> *href here should match the HandleFunc path below 
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "frontend/src/index.html")
	})

	log.Fatal(http.ListenAndServe(":8000", mux))
}