package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

const FIVE = 5

type TestMessage struct {
	Text string `json:"text"`
}

type UserLeaderboard struct {
	Id         int     `json:"id"`
	Rank       int     `json:"rank"`
	Name       string  `json:"name"`
	SumDonated float64 `json:"sum_donated"`
}

func getLeaderboardTopFive() [FIVE]UserLeaderboard { // after i have the db this should be created from db (sorted desc by sum_donated)
	var leaderboardTopFive [FIVE]UserLeaderboard
	leaderboardTopFive[0] = UserLeaderboard{1, 1, "Name1", 1}
	leaderboardTopFive[1] = UserLeaderboard{2, 2, "Name2", 2}
	leaderboardTopFive[2] = UserLeaderboard{3, 3, "Name3", 3}
	leaderboardTopFive[3] = UserLeaderboard{4, 4, "Name4", 4}
	leaderboardTopFive[4] = UserLeaderboard{5, 5, "Name5", 5}

	return leaderboardTopFive

}

func homePage(writer http.ResponseWriter, r *http.Request) { //http.ResponseWriter = send data to HTTP client; http.Request = represent the client

	var leaderboardTopFive [FIVE]UserLeaderboard = getLeaderboardTopFive()

	html_template := template.Must(template.ParseFiles("../frontend/index.html"))
	log.Println(html_template.Execute(writer, leaderboardTopFive))

}

func testPage(writer http.ResponseWriter, r *http.Request) { //http.ResponseWriter = send data to HTTP client; http.Request = represent the client
	fmt.Fprintf(writer, "This is the TEST page from GO")
}

// func apiHandler(w http.ResponseWriter, r *http.Request) {

// 	msg := TestMessage{Text: "Hello from Go backend!"}
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(msg)
// }

func main() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/test", testPage)

	fmt.Println("Server is running on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
