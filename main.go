package main

import (
	"github.com/EthanG78/GoBall/utils"
	"html/template"
	"log"
	"net/http"
	"strings"
)

var PlayersMap = map[string]int{"0": 0}

func index(w http.ResponseWriter, r *http.Request) {

	var first string
	var last string
	var stats utils.PlayerStats

	r.ParseForm()

	for key, value := range r.Form {
		if key == "first_name" {
			first = strings.Join(value, " ")
		} else if key == "last_name" {
			last = strings.Join(value, " ")
		}
	}

	stats = utils.FetchPlayerStats(first, last)

	if utils.ContainsKey(stats.Name, PlayersMap) {
		PlayersMap[stats.Name]++
	} else if !utils.ContainsKey(stats.Name, PlayersMap) {
		PlayersMap[stats.Name] = 1
	}

	utils.PrintMap(PlayersMap)

	t, err := template.ParseFiles("static/index.html")
	if err != nil {
		log.Fatalf("Error parsing template: %v", err)
	}

	t.Execute(w, stats)
}

func main() {
	http.HandleFunc("/", index)
	log.Println("Now Serving on port 8080")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
