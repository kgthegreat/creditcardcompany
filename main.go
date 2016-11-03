package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {

	var (
		addr string = ":4000"
	)

	flag.StringVar(&addr, "addr", ":4000", "")
	flag.Parse()

	server := NewServer(addr)
	StartServer(server)
}

func NewServer(addr string) *http.Server {
	// Setup router
	//	initRouting()
	initStaticRouting()
	
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/robots.txt", robotsHandler)
	//http.HandleFunc("/compare", comparehandler)
	// Create and start server
	return &http.Server{
		Addr:    addr,
	}
}

func StartServer(server *http.Server) {
	log.Println("Starting server")
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalln("Error: %v", err)
	}
}

func robotsHandler(w http.ResponseWriter, req *http.Request) {
	
	var empty []string
	renderTemplate(w, "robots", empty)

}

func indexHandler(w http.ResponseWriter, req *http.Request) {
	
	siteName := "https://www.moneyhero.com.hk"
	url := getUrl(siteName)
	resp := getResponse(url)
	varmap := getVarmap(resp)
	defer resp.Body.Close()

	renderTemplate(w, "index", varmap)

}

