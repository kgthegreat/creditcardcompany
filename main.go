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
	http.HandleFunc("/hongkong", hkCCHandler)
	http.HandleFunc("/singapore", sgCCHandler)
	http.HandleFunc("/taiwan", twCCHandler)
	http.HandleFunc("/portugal", ptCCHandler)
	http.HandleFunc("/hackathon", hackCCHandler)
	http.HandleFunc("/robots.txt", robotsHandler)
	//http.HandleFunc("/compare", comparehandler)
	// Create and start server
	return &http.Server{
		Addr: addr,
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

	var empty []string
	renderTemplate(w, "index", empty)

}

func hkCCHandler(w http.ResponseWriter, req *http.Request) {

	siteName := "https://www.moneyhero.com.hk"
	url := getUrl(siteName)
	resp := getResponse(url)
	varmap := getVarmap(resp)
	defer resp.Body.Close()

	renderTemplate(w, "cc", varmap)

}

func sgCCHandler(w http.ResponseWriter, req *http.Request) {

	siteName := "https://www.singsaver.com.sg"
	url := getUrl(siteName)
	resp := getResponse(url)
	varmap := getVarmap(resp)
	defer resp.Body.Close()

	renderTemplate(w, "cc", varmap)

}

func twCCHandler(w http.ResponseWriter, req *http.Request) {

	siteName := "https://www.money101.com.tw"
	url := getUrl(siteName)
	resp := getResponse(url)
	varmap := getVarmap(resp)
	defer resp.Body.Close()

	renderTemplate(w, "cc", varmap)

}

func ptCCHandler(w http.ResponseWriter, req *http.Request) {

	siteName := "https://www.comparaja.pt"
	url := getUrl(siteName)
	resp := getResponse(url)
	varmap := getVarmap(resp)
	defer resp.Body.Close()

	renderTemplate(w, "cc", varmap)

}

func hackCCHandler(w http.ResponseWriter, req *http.Request) {

	siteName := "https://hackathon.compareglobal.co.uk"
	url := getUrl(siteName)
	resp := getResponse(url)
	varmap := getVarmap(resp)
	defer resp.Body.Close()

	renderTemplate(w, "cc", varmap)

}
