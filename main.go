package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

// templateData provides template parameters.
type templateData struct {
	Service  string
	Revision string
}

type Player struct {
	Name  string
	Team  string
	Price int
}

// Variables used to generate the HTML page.
var (
	data templateData
	tmpl *template.Template
)

func main() {
	// Initialize template parameters.
	service := os.Getenv("K_SERVICE")
	if service == "" {
		service = "???"
	}

	revision := os.Getenv("K_REVISION")
	if revision == "" {
		revision = "???"
	}

	// Prepare template for execution.
	tmpl = template.Must(template.ParseFiles("index.html"))
	data = templateData{
		Service:  service,
		Revision: revision,
	}

	// Define HTTP server.
	http.HandleFunc("/", helloRunHandler)

	fs := http.FileServer(http.Dir("./assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	// PORT environment variable is provided by Cloud Run.
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Print("Hello from Cloud Run! The container started successfully and is listening for HTTP requests on $PORT")
	log.Printf("Listening on port %s", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}

// helloRunHandler responds to requests by rendering an HTML page.
func helloRunHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	squad := make(map[string][]Player)
	squad["Fullbacks"] = []Player{
		{Name: "Trent Alexander-Arnold", Team: "Liverpool", Price: 7},
		{Name: "Andrew Robertson", Team: "Liverpool", Price: 7},
	}
	squad["CentreWingers"] = []Player{
		{Name: "Mohamed Salah", Team: "Liverpool", Price: 12},
		{Name: "Sadio Mane", Team: "Liverpool", Price: 12},
	}
	squad["FiveEights"] = []Player{}
	squad["Halfbacks"] = []Player{}
	squad["FrontRowers"] = []Player{}
	squad["Hookers"] = []Player{}
	squad["SecondRowers"] = []Player{}
	enableCors(&w)
	jsonData, err := json.Marshal(squad)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Print(jsonData)
	w.Write(jsonData)
	// tmpl.Execute(w, data)

}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}
