package main

import (
	"flag"
	"log"
	"net/http"
	"fmt"
	"encoding/json"
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

	cssHandler := http.FileServer(http.Dir("./static/css/"))
	jsHandler := http.FileServer(http.Dir("./static/js/"))
	imgHandler := http.FileServer(http.Dir("./static/img/"))
	fontsHandler := http.FileServer(http.Dir("./static/fonts/"))

	http.Handle("/css/", http.StripPrefix("/css/", cssHandler))
	http.Handle("/fonts/", http.StripPrefix("/fonts/", fontsHandler))
	http.Handle("/js/", http.StripPrefix("/js/", jsHandler))
	http.Handle("/img/", http.StripPrefix("/img/", imgHandler))

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
	
	log.Println("Inside ih")


	siteName := "https://www.moneyhero.com.hk"
	category := "airmiles"
	language := "en"
	pageSize := "6"

	url := fmt.Sprintf("%s/api/credit-card/v2/cards/%s?&lang=%s&pageSize=%s",siteName,category,language,pageSize)

	// Build the request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return
	}

	// For control over HTTP client headers,
	// redirect policy, and other settings,
	// create a Client
	// A Client is an HTTP client
	client := &http.Client{}

	// Send the request via a client
	// Do sends an HTTP request and
	// returns an HTTP response
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return
	}

	// Callers should close resp.Body
	// when done reading from it
	// Defer the closing of the body
	defer resp.Body.Close()

	// Fill the record with the data from the JSON
	var creditCards CreditCards

	// Use json.Decode for reading streams of JSON data
	if err := json.NewDecoder(resp.Body).Decode(&creditCards); err != nil {
		log.Println(err)
	}

	fmt.Println("Name = ", creditCards.FeaturedCards[0].Name)
//	FeaturedCards := 
	varmap := map[string]interface{}{
		"FeaturedCards": creditCards.FeaturedCards,
		"AllCards": creditCards.Cards,
	}
//	var empty []string
	renderTemplate(w, "index", varmap)

}
