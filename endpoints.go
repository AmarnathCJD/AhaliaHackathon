package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func SignUpCharityW(w http.ResponseWriter, r *http.Request) {
	var charity Charity
	err := json.NewDecoder(r.Body).Decode(&charity)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	charity.EType = "charity"

	err = SignUpCharity(charity)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func SignUpRestaurantW(w http.ResponseWriter, r *http.Request) {
	var restaurant Restaurant
	err := json.NewDecoder(r.Body).Decode(&restaurant)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	restaurant.EType = "restaurant"

	err = SignUpRestaurant(restaurant)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func Login(w http.ResponseWriter, r *http.Request) {
	cors(w, r)
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	} else if r.Method != "POST" {
		http.Error(w, "{\"error\": \"method not allowed\"}", http.StatusMethodNotAllowed)
		return
	}
	var user User
	if r.Header.Get("Content-Type") == "application/json" {
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	} else {
		r.ParseForm()
		user.Username = r.FormValue("username")
		user.Password = r.FormValue("password")
	}

	isValid, _type, token, err := SignIn(user.Username, user.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if !isValid {
		http.Error(w, "invalid credentials", http.StatusUnauthorized)
		return
	}

	var resp = struct {
		Valid bool   `json:"valid"`
		Type  string `json:"type"`
		Token string `json:"token"`
	}{
		Valid: isValid,
		Type:  _type,
		Token: token,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func GetTokenUsernameW(w http.ResponseWriter, r *http.Request) {
	cors(w, r)
	var token struct {
		Token string `json:"token,omitempty"`
	}

	if r.Method == "GET" {
		r.ParseForm()
		token.Token = r.FormValue("token")
	} else {
		err := json.NewDecoder(r.Body).Decode(&token)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}

	username, userType, err := checkSessionToken(token.Token)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Write([]byte(fmt.Sprintf(`{"username": "%s", "user_type": "%s"}`, username, userType)))
}

func SearchForWasteSourcesW(w http.ResponseWriter, r *http.Request) {
	cors(w, r)
	var src = struct {
		Location string `json:"location,omitempty"`
		Name     string `json:"name,omitempty"`
		Filter   string `json:"filter,omitempty"`
	}{
		Location: r.FormValue("location"),
		Name:     r.FormValue("name"),
		Filter:   r.FormValue("filter"),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(SearchForWasteSources(src.Location, src.Name, src.Filter))
}

func GetWasteSourceW(w http.ResponseWriter, r *http.Request) {
	cors(w, r)
	var src struct {
		Username string `json:"username,omitempty"`
	}

	if r.Method == "GET" {
		r.ParseForm()
		src.Username = r.FormValue("username")
	} else {
		err := json.NewDecoder(r.Body).Decode(&src)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	for _, charity := range Charities {
		if charity.Username == src.Username {
			json.NewEncoder(w).Encode(charity)
			return
		}
	}

	for _, restaurant := range Restaurants {
		if restaurant.Username == src.Username {
			json.NewEncoder(w).Encode(restaurant)
			return
		}
	}

	for _, buisness := range Buisnesses {
		if buisness.Username == src.Username {
			json.NewEncoder(w).Encode(buisness)
			return
		}
	}

	http.Error(w, "not found", http.StatusNotFound)
}

// Function to get the GeoData from the given Latitude and Longitude.
func GetGeoIPW(w http.ResponseWriter, r *http.Request) {
	cors(w, r)
	var ip struct {
		IP string `json:"ip,omitempty"`
	}

	if r.Method == "GET" {
		r.ParseForm()
		ip.IP = r.FormValue("ip")
	} else {
		err := json.NewDecoder(r.Body).Decode(&ip)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	lat, long, data := getGeoDecode(ip.IP)
	json.NewEncoder(w).Encode(struct {
		Lat  string `json:"lat"`
		Long string `json:"long"`
		Data string `json:"data"`
	}{
		Lat:  lat,
		Long: long,
		Data: data,
	})
}

func RequestWasteFood(w http.ResponseWriter, r *http.Request) {
	cors(w, r)
	var order Order
	if r.Method == "GET" {
		r.ParseForm()
		order.Orderee = r.FormValue("orderee")
		order.Orderer = r.FormValue("orderer")
		order_ID, _ := strconv.ParseInt(r.FormValue("id"), 10, 64)
		order.ID = int(order_ID)
		order.Product = r.FormValue("product")
		order_Quantity, _ := strconv.ParseInt(r.FormValue("quantity"), 10, 64)
		order.Quantity = int(order_Quantity)
	} else {
		err := json.NewDecoder(r.Body).Decode(&order)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}

	err := RequestWaste(order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func GetPendingOrders(w http.ResponseWriter, r *http.Request) {
	cors(w, r)
	var username struct {
		Username string `json:"username,omitempty"`
	}

	if r.Method == "GET" {
		r.ParseForm()
		username.Username = r.FormValue("username")
	} else {
		err := json.NewDecoder(r.Body).Decode(&username)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")

	if username.Username == "" {
		var orders []Order
		for _, charity := range Charities {
			orders = append(orders, charity.GetOrders()...)
		}

		for _, restaurant := range Restaurants {
			orders = append(orders, restaurant.GetOrders()...)
		}

		// remove duplicates and sort by time
		orders = removeDuplicates(orders)
		orders = sortOrders(orders)

		json.NewEncoder(w).Encode(orders)
		return
	}

	for _, charity := range Charities {
		if charity.Username == username.Username {
			json.NewEncoder(w).Encode(charity.GetOrders())
			return
		}
	}

	for _, restaurant := range Restaurants {
		if restaurant.Username == username.Username {
			json.NewEncoder(w).Encode(restaurant.GetOrders())
			return
		}
	}

	http.Error(w, "not found", http.StatusNotFound)
}

func AcceptOrder(w http.ResponseWriter, r *http.Request) {
	cors(w, r)
	var opt struct {
		Username string `json:"username,omitempty"`
		ID       int    `json:"id,omitempty"`
	}

	if r.Method == "GET" {
		r.ParseForm()
		opt.Username = r.FormValue("username")
		opt_ID, _ := strconv.ParseInt(r.FormValue("id"), 10, 64)
		opt.ID = int(opt_ID)
	} else {
		err := json.NewDecoder(r.Body).Decode(&opt)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}

	for _, charity := range Charities {
		if charity.Username == opt.Username {
			err := charity.AcceptOrder(opt.ID)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			w.WriteHeader(http.StatusOK)
			return
		}
	}

	for _, restaurant := range Restaurants {
		if restaurant.Username == opt.Username {
			err := restaurant.AcceptOrder(opt.ID)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			w.WriteHeader(http.StatusOK)
			return
		}
	}

	for _, buisness := range Buisnesses {
		if buisness.Username == opt.Username {
			err := buisness.AcceptOrder(opt.ID)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			w.WriteHeader(http.StatusOK)
			return
		}
	}

	http.Error(w, "not found", http.StatusNotFound)
}

func DeclineOrder(w http.ResponseWriter, r *http.Request) {
	cors(w, r)
	var opt struct {
		Username string `json:"username,omitempty"`
		ID       int    `json:"id,omitempty"`
	}

	if r.Method == "GET" {
		r.ParseForm()
		opt.Username = r.FormValue("username")
		optID, _ := strconv.ParseInt(r.FormValue("id"), 10, 64)
		opt.ID = int(optID)
	} else {
		err := json.NewDecoder(r.Body).Decode(&opt)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}

	for _, charity := range Charities {
		if charity.Username == opt.Username {
			err := charity.DeleteOrder(opt.ID)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			w.WriteHeader(http.StatusOK)
			return
		}
	}

	for _, restaurant := range Restaurants {
		if restaurant.Username == opt.Username {
			err := restaurant.DeleteOrder(opt.ID)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			w.WriteHeader(http.StatusOK)
			return
		}
	}

	for _, buisness := range Buisnesses {
		if buisness.Username == opt.Username {
			err := buisness.DeleteOrder(opt.ID)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			w.WriteHeader(http.StatusOK)
			return
		}
	}

	http.Error(w, "not found", http.StatusNotFound)
}

func AddEntry(w http.ResponseWriter, r *http.Request) {
	cors(w, r)
	var entry struct {
		Username string `json:"username,omitempty"`
		Name     string `json:"name,omitempty"`
		Quantity int    `json:"quantity,omitempty"`
		Type     string `json:"type,omitempty"`
	}

	if r.Method == "GET" || r.Method == "POST" {
		r.ParseForm()
		entry.Username = r.FormValue("username")
		entry.Name = r.FormValue("name")
		entryQuantity, _ := strconv.ParseInt(r.FormValue("quantity"), 10, 64)
		entry.Quantity = int(entryQuantity)
		entry.Type = r.FormValue("type")
	} else {
		err := json.NewDecoder(r.Body).Decode(&entry)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}

	for _, charity := range Buisnesses {
		if charity.Username == entry.Username {
			err := charity.AddEntry(entry.Name, entry.Quantity, entry.Type)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			w.WriteHeader(http.StatusOK)
			return
		}
	}

	for _, restaurant := range Restaurants {
		if restaurant.Username == entry.Username {
			err := restaurant.AddEntry(entry.Name, entry.Quantity, entry.Type)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			w.WriteHeader(http.StatusOK)
			return
		}
	}
}

func removeDuplicates(orders []Order) []Order {
	keys := make(map[int]bool)
	list := []Order{}
	for _, entry := range orders {
		if _, value := keys[entry.ID]; !value {
			keys[entry.ID] = true
			list = append(list, entry)
		}
	}
	return list
}

func sortOrders(orders []Order) []Order {
	for i := 0; i < len(orders); i++ {
		for j := 0; j < len(orders)-1; j++ {
			if orders[j].DateTime < orders[j+1].DateTime {
				orders[j], orders[j+1] = orders[j+1], orders[j]
			}
		}
	}
	return orders
}
