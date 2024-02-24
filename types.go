package main

const MONGODB_URI = "mongodb+srv://vwaste:vwaste@cluster0.1vyfegeeee.mongodb.net/?retryWrites=true&w=majority"

type Order struct {
	ID        int    `json:"id,omitempty"`
	Product   string `json:"product,omitempty"`
	Quantity  int    `json:"quantity,omitempty"`
	Approved  bool   `json:"approved"`
	DateTime  int64  `json:"date_time,omitempty"`
	Orderer   string `json:"orderer,omitempty"`
	Orderee   string `json:"orderee,omitempty"`
	OrderType string `json:"order_type,omitempty"`
	EType     string `json:"e_type,omitempty"` // placeholder for type Interface
}

type Waste struct {
	Product  string `json:"product,omitempty"`
	Quantity int    `json:"quantity,omitempty"`
	Edible   bool   `json:"edible,omitempty"`
	Family   string `json:"family,omitempty"` // veg or non-veg
}

type Charity struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Location    string `json:"location,omitempty"`
	Username    string `json:"username,omitempty"`
	Password    string `json:"password,omitempty"`

	Orders []Order `json:"orders,omitempty"`
	EType  string  `json:"e_type,omitempty"` // placeholder for type Interface
}

type Review struct {
	RestaurantUsername string  `json:"restaurant_username,omitempty"`
	Content            string  `json:"content,omitempty"`
	Rating             float64 `json:"rating,omitempty"`
}

type Point struct {
	Rating              float64 `json:"rating,omitempty"`
	LeaderboardPosition int     `json:"leaderboard_position,omitempty"`
}

type Restaurant struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Location    string `json:"location,omitempty"`
	Username    string `json:"username,omitempty"`
	Password    string `json:"password,omitempty"`

	Reviews []Review `json:"reviews,omitempty"`
	Others  []Order  `json:"others,omitempty"`
	Point   Point    `json:"point,omitempty"`
	Waste   []Waste  `json:"waste,omitempty"`
	EType   string   `json:"e_type,omitempty"` // placeholder for type Interface
}

type Buisness struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Location    string `json:"location,omitempty"`
	Username    string `json:"username,omitempty"`
	Password    string `json:"password,omitempty"`

	AISubscription bool    `json:"ai_subscription,omitempty"`
	Orders         []Order `json:"orders,omitempty"`
	Waste          []Waste `json:"waste,omitempty"`
}

type User struct {
	Username     string `json:"username,omitempty"`
	Password     string `json:"password,omitempty"`
	SessionToken string `json:"session_token,omitempty"`
	UserType     string `json:"user_type,omitempty"`
}

type UserList struct {
	Users []User `json:"users,omitempty"`
}
