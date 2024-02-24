package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Hello, World!")

	http.ListenAndServe(":8080", nil)
}

func init() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		cors(w, r)
		fmt.Fprint(w, "Hello, World!")
	})

	// ---- Auth Endpoints ----
	http.HandleFunc("/api/signup/charity", SignUpCharityW)
	http.HandleFunc("/api/signup/restaurant", SignUpRestaurantW) //
	http.HandleFunc("/api/login", Login)                         // {username: "username", password: "password"}
	http.HandleFunc("/api/getsession", GetTokenUsernameW)        // {token: "token"}

	// ---- Waste Sources Data ----
	http.HandleFunc("/api/search", SearchForWasteSourcesW) // {location: "location", name: "name", filter: "filter"}

	// ---- Waste Sources Endpoints ----
	http.HandleFunc("/api/get", GetWasteSourceW) // {username: "username"}

	// ---- MISCELLANEOUS ----
	http.HandleFunc("/api/ip", GetGeoIPW) // {ip: "ip"}

	http.HandleFunc("/api/getorders", GetPendingOrders) // {username: "username"}
	http.HandleFunc("/api/neworder", RequestWasteFood)  // {username: "username", order: Order{}}
	http.HandleFunc("/api/acceptorder", AcceptOrder)    // {username: "username", order: Order{}}
	http.HandleFunc("/api/declineorder", DeclineOrder)  // {username: "username", order: Order{}}

	http.HandleFunc("/api/addentry", AddEntry) // {username: "username", entry: Entry{}}

}

// cors Middleware for Cross Domain res requests.
func cors(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", req.Header.Get("Origin"))
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}
