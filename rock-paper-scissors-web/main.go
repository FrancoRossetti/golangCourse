package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"myapp/rps"
	"net/http"
	"strconv"
	"text/template"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index.html")
}

var sum = rps.Summary{
	ComputerWins: 0,
	PlayerWins:   0,
	Draws:        0,
}

type RandomData struct {
	RandomMessage string `json:"random_message"`
}

func playRound(w http.ResponseWriter, r *http.Request) {
	playerChoice, _ := strconv.Atoi(r.URL.Query().Get("c"))
	result, summary := rps.PlayRound(sum, playerChoice)
	sum = summary
	out, err := json.MarshalIndent(result, "", "   ")
	if err != nil {
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}
func main() {

	http.HandleFunc("/", homePage)
	http.HandleFunc("/play", playRound)
	http.HandleFunc("/getStatistics", getStatistics)
	http.HandleFunc("/getRandomData", getRandomData)

	log.Println("Starting web server on port 8080")

	http.ListenAndServe(":8080", nil)
}

func renderTemplate(w http.ResponseWriter, page string) {
	t, err := template.ParseFiles("index.html")
	if err != nil {
		log.Println(err)
		return
	}
	err = t.Execute(w, nil)
	if err != nil {
		log.Println(err)
		return
	}
}

func getStatistics(w http.ResponseWriter, r *http.Request) {
	out, err := json.MarshalIndent(sum, "", "   ")
	if err != nil {
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

func getRandomData(w http.ResponseWriter, r *http.Request) {
	playerChoice, _ := strconv.Atoi(r.URL.Query().Get("number"))
	var apiurl = fmt.Sprintf("http://numbersapi.com/%d/trivia", playerChoice)
	response, err := http.Get(apiurl)
	if err != nil {
		log.Println(err)
		return
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
		random := RandomData{
			RandomMessage: string(data),
		}
		out, err := json.MarshalIndent(random, "", "   ")
		if err != nil {
			log.Println(err)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(out)
	}
}
