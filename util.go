package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"encoding/json"
	"fmt"
)

var templates *template.Template

func init() {
	filenames := []string{}
	err := filepath.Walk("templates", func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() && filepath.Ext(path) == ".html" {
			filenames = append(filenames, path)
		}

		return nil
	})
	if err != nil {
		log.Fatalln(err)
	}

	if len(filenames) == 0 {
		return
	}


	templates, err = template.New("").Funcs(template.FuncMap{
		"convertMinusOne": func(inp string) string {
			if inp == "-1" {
				return "No Fees"
			} else {
				return inp
			}

		},
		"truncateLongText": func(inp string) string {
			return inp[0:25] + " ..."
		},
	}).ParseFiles(filenames...)
	if err != nil {
		log.Fatalln(err)
	}

}

func initStaticRouting() {
	cssHandler := http.FileServer(http.Dir("./static/css/"))
	jsHandler := http.FileServer(http.Dir("./static/js/"))
	imgHandler := http.FileServer(http.Dir("./static/img/"))
	fontsHandler := http.FileServer(http.Dir("./static/fonts/"))

	http.Handle("/css/", http.StripPrefix("/css/", cssHandler))
	http.Handle("/fonts/", http.StripPrefix("/fonts/", fontsHandler))
	http.Handle("/js/", http.StripPrefix("/js/", jsHandler))
	http.Handle("/img/", http.StripPrefix("/img/", imgHandler))
}

func renderTemplate(w http.ResponseWriter, tmpl string, vars interface{}) {
	
	err := templates.ExecuteTemplate(w, tmpl+".html", vars)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func convertMinusOne(inp string) string {
	if inp == "-1" {
		return "No Fees"
	} else {
		return inp
	}

}

func getResponse(url string) *http.Response {
	// Build the request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
	//	return
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
	//	return
	}
	return resp
}

func getVarmap(resp *http.Response) map[string]interface{} {
	// Fill the record with the data from the JSON
	var creditCards CreditCards

	// Use json.Decode for reading streams of JSON data
	if err := json.NewDecoder(resp.Body).Decode(&creditCards); err != nil {
		log.Println(err)
	}

	varmap := map[string]interface{}{
		"FeaturedCards": creditCards.FeaturedCards,
		"AllCards": creditCards.Cards,
	}

	return varmap

}

func getUrl(siteName string) string {
	category := "airmiles"
	language := "en"
	pageSize := "6"

	url := fmt.Sprintf("%s/api/credit-card/v2/cards/%s?&lang=%s&pageSize=%s",siteName,category,language,pageSize)
	return url
}

