package main

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	opts         = options.Update().SetUpsert(true)
	GlobalOrders []Order
)

func __load_All_ENTITIES() {
	var charities []Charity
	var restaurants []Restaurant
	var buisnesses []Buisness
	var users UserList

	collection := DB.Database("vwaste").Collection("charities")

	cursor, _ := collection.Find(context.TODO(), bson.M{})
	cursor.All(context.TODO(), &charities)

	collection = DB.Database("vwaste").Collection("restaurants")

	cursor, _ = collection.Find(context.TODO(), bson.M{})
	cursor.All(context.TODO(), &restaurants)

	collection = DB.Database("vwaste").Collection("users")

	cursor, _ = collection.Find(context.TODO(), bson.M{})
	cursor.All(context.TODO(), &users)

	Charities = charities
	Restaurants = restaurants
	Users = users

	collection = DB.Database("vwaste").Collection("buisnesses")
	cursor, _ = collection.Find(context.TODO(), bson.M{})
	cursor.All(context.TODO(), &buisnesses)

	Buisnesses = buisnesses
}

func __drop_ALL() {
	DB.Database("vwaste").Collection("charities").Drop(context.Background())
	DB.Database("vwaste").Collection("restaurants").Drop(context.Background())
	DB.Database("vwaste").Collection("users").Drop(context.Background())
	DB.Database("vwaste").Collection("buisnesses").Drop(context.Background())
}

var (
	Charities   []Charity
	Restaurants []Restaurant
	Buisnesses  []Buisness
	Users       UserList
	DB          *mongo.Client
)

func init() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(MONGODB_URI))
	if err != nil {
		panic("Error connecting to MongoDB")
	}

	DB = client
	__load_All_ENTITIES()
	__load_globalOrders()
}

// ---------- DB METHODS ----------

func (c *Charity) Save() error {
	collection := DB.Database("vwaste").Collection("charities")
	_, err := collection.UpdateOne(context.Background(), bson.M{"username": c.Username}, bson.M{"$set": c}, opts)

	for i, charity := range Charities {
		if charity.Username == c.Username {
			Charities[i] = *c
			break
		}
	}
	return err
}

func (r *Restaurant) Save() error {
	collection := DB.Database("vwaste").Collection("restaurants")
	_, err := collection.UpdateOne(context.Background(), bson.M{"username": r.Username}, bson.M{"$set": r}, opts)

	for i, restaurant := range Restaurants {
		if restaurant.Username == r.Username {
			Restaurants[i] = *r
			break
		}
	}
	return err
}

func (b *Buisness) Save() error {
	collection := DB.Database("vwaste").Collection("buisnesses")
	_, err := collection.UpdateOne(context.Background(), bson.M{"username": b.Username}, bson.M{"$set": b}, opts)

	for i, buisness := range Buisnesses {
		if buisness.Username == b.Username {
			Buisnesses[i] = *b
			break
		}
	}

	return err
}

func (u *User) Save() error {
	collection := DB.Database("vwaste").Collection("users")
	_, err := collection.UpdateOne(context.Background(), bson.M{"username": u.Username}, bson.M{"$set": u}, opts)

	for i, user := range Users.Users {
		if user.Username == u.Username {
			Users.Users[i] = *u
			break
		}
	}
	return err
}

func (c *Charity) Delete() error {
	collection := DB.Database("vwaste").Collection("charities")
	_, err := collection.DeleteOne(context.Background(), bson.M{"username": c.Username})

	for i, charity := range Charities {
		if charity.Username == c.Username {
			Charities = append(Charities[:i], Charities[i+1:]...)
			break
		}
	}
	return err
}

func (r *Restaurant) Delete() error {
	collection := DB.Database("vwaste").Collection("restaurants")
	_, err := collection.DeleteOne(context.Background(), bson.M{"username": r.Username})

	for i, restaurant := range Restaurants {
		if restaurant.Username == r.Username {
			Restaurants = append(Restaurants[:i], Restaurants[i+1:]...)
			break
		}
	}
	return err
}

func (b *Buisness) Delete() error {
	collection := DB.Database("vwaste").Collection("buisnesses")
	_, err := collection.DeleteOne(context.Background(), bson.M{"username": b.Username})

	for i, buisness := range Buisnesses {
		if buisness.Username == b.Username {
			Buisnesses = append(Buisnesses[:i], Buisnesses[i+1:]...)
			break
		}
	}

	return err
}

func (u *User) Delete() error {
	collection := DB.Database("vwaste").Collection("users")
	_, err := collection.DeleteOne(context.Background(), bson.M{"username": u.Username})

	for i, user := range Users.Users {
		if user.Username == u.Username {
			Users.Users = append(Users.Users[:i], Users.Users[i+1:]...)
			break
		}
	}
	return err
}

func (c *Charity) AddOrder(o Order) error {
	o.DateTime = time.Now().Unix()
	o.Approved = false
	collection := DB.Database("vwaste").Collection("charities")
	_, err := collection.UpdateOne(context.Background(), bson.M{"username": c.Username}, bson.M{"$push": bson.M{"orders": o}}, opts)

	if err != nil {
		var charity Charity
		collection.FindOne(context.Background(), bson.M{"username": c.Username}).Decode(&charity)
		charity.Orders = append(charity.Orders, o)

		for i, charity := range Charities {
			if charity.Username == c.Username {
				Charities[i].Orders = append(Charities[i].Orders, o)
				break
			}
		}

		collection.UpdateOne(context.Background(), bson.M{"username": c.Username}, bson.M{"$set": charity}, opts)
		return err
	}

	for i, charity := range Charities {
		if charity.Username == c.Username {
			Charities[i].Orders = append(Charities[i].Orders, o)
			break
		}
	}
	return err
}

func (r *Restaurant) AddOrder(o Order) error {
	o.DateTime = time.Now().Unix()
	o.Approved = false
	collection := DB.Database("vwaste").Collection("restaurants")
	_, err := collection.UpdateOne(context.Background(), bson.M{"username": r.Username}, bson.M{"$push": bson.M{"others": o}}, opts)
	if err != nil {
		var restaurant Restaurant
		collection.FindOne(context.Background(), bson.M{"username": r.Username}).Decode(&restaurant)
		restaurant.Others = append(restaurant.Others, o)

		for i, restaurant := range Restaurants {
			if restaurant.Username == r.Username {
				Restaurants[i].Others = append(Restaurants[i].Others, o)
				break
			}
		}

		collection.UpdateOne(context.Background(), bson.M{"username": r.Username}, bson.M{"$set": restaurant}, opts)
		return err
	}

	for i, restaurant := range Restaurants {
		if restaurant.Username == r.Username {
			Restaurants[i].Others = append(Restaurants[i].Others, o)
			break
		}
	}
	return err
}

func (r *Restaurant) AddReview(rev Review) error {
	collection := DB.Database("vwaste").Collection("restaurants")
	_, err := collection.UpdateOne(context.Background(), bson.M{"username": r.Username}, bson.M{"$push": bson.M{"reviews": rev}}, opts)

	if err != nil {
		var restaurant Restaurant
		collection.FindOne(context.Background(), bson.M{"username": r.Username}).Decode(&restaurant)
		restaurant.Reviews = append(restaurant.Reviews, rev)

		for i, restaurant := range Restaurants {
			if restaurant.Username == r.Username {
				Restaurants[i].Reviews = append(Restaurants[i].Reviews, rev)
				break
			}
		}

		collection.UpdateOne(context.Background(), bson.M{"username": r.Username}, bson.M{"$set": restaurant}, opts)
		return err
	}

	for i, restaurant := range Restaurants {
		if restaurant.Username == r.Username {
			Restaurants[i].Reviews = append(Restaurants[i].Reviews, rev)
			break
		}
	}
	return err
}

func (b *Buisness) AddOrder(o Order) error {
	o.DateTime = time.Now().Unix()
	o.Approved = false
	collection := DB.Database("vwaste").Collection("buisnesses")
	_, err := collection.UpdateOne(context.Background(), bson.M{"username": b.Username}, bson.M{"$push": bson.M{"orders": o}}, opts)

	if err != nil {
		var buisness Buisness
		collection.FindOne(context.Background(), bson.M{"username": b.Username}).Decode(&buisness)
		buisness.Orders = append(buisness.Orders, o)

		for i, buisness := range Buisnesses {
			if buisness.Username == b.Username {
				Buisnesses[i].Orders = append(Buisnesses[i].Orders, o)
				break
			}
		}

		collection.UpdateOne(context.Background(), bson.M{"username": b.Username}, bson.M{"$set": buisness}, opts)
		return err
	}

	for i, buisness := range Buisnesses {
		if buisness.Username == b.Username {
			Buisnesses[i].Orders = append(Buisnesses[i].Orders, o)
			break
		}
	}
	return err
}

func (c *Charity) AcceptOrder(id int) error {
	collection := DB.Database("vwaste").Collection("charities")
	for _, char := range Charities {
		for i, order := range char.Orders {
			if order.ID == id {
				char.Orders[i].Approved = true
				_, err := collection.UpdateOne(context.Background(), bson.M{"username": c.Username}, bson.M{"$set": bson.M{"orders": char.Orders}}, opts)
				return err
			}
		}
	}
	return nil
}

func (r *Restaurant) AcceptOrder(id int) error {
	collection := DB.Database("vwaste").Collection("restaurants")
	for i, order := range r.Others {
		if order.ID == id {
			r.Others[i].Approved = true
			_, err := collection.UpdateOne(context.Background(), bson.M{"username": r.Username}, bson.M{"$set": bson.M{"others": r.Others}}, opts)
			return err
		}
	}
	return nil
}

func (b *Buisness) AcceptOrder(id int) error {
	collection := DB.Database("vwaste").Collection("buisnesses")
	for i, order := range b.Orders {
		if order.ID == id {
			b.Orders[i].Approved = true
			_, err := collection.UpdateOne(context.Background(), bson.M{"username": b.Username}, bson.M{"$set": bson.M{"orders": b.Orders}}, opts)
			return err
		}
	}
	return nil
}

func (c *Charity) DeleteOrder(id int) error {
	collection := DB.Database("vwaste").Collection("charities")
	for i, order := range c.Orders {
		if order.ID == id {
			c.Orders = append(c.Orders[:i], c.Orders[i+1:]...)
			_, err := collection.UpdateOne(context.Background(), bson.M{"username": c.Username}, bson.M{"$set": bson.M{"orders": c.Orders}}, opts)
			return err
		}
	}
	return nil
}

func (r *Restaurant) DeleteOrder(id int) error {
	collection := DB.Database("vwaste").Collection("restaurants")
	for i, order := range r.Others {
		if order.ID == id {
			if len(r.Others) == 1 {
				r.Others = []Order{}
			} else {
				r.Others = append(r.Others[:i], r.Others[i+1:]...)
			}
			_, err := collection.UpdateOne(context.Background(), bson.M{"username": r.Username}, bson.M{"$set": bson.M{"others": r.Others}}, opts)
			return err
		}
	}
	return nil
}

func (b *Buisness) DeleteOrder(id int) error {
	collection := DB.Database("vwaste").Collection("buisnesses")
	for i, order := range b.Orders {
		if order.ID == id {
			Buisnesses[i].Orders = append(Buisnesses[i].Orders[:i], Buisnesses[i].Orders[i+1:]...)
			_, err := collection.UpdateOne(context.Background(), bson.M{"username": b.Username}, bson.M{"$set": bson.M{"orders": b.Orders}}, opts)
			return err
		}
	}
	return nil
}

func (u *User) AddUser() error {
	collection := DB.Database("vwaste").Collection("users")
	_, err := collection.UpdateOne(context.Background(), bson.M{"username": u.Username}, bson.M{"$set": u}, opts)

	Users.Users = append(Users.Users, *u)
	return err
}

func (c *Charity) GetOrders() []Order {
	var orders []Order = c.Orders
	for _, order := range Restaurants {
		for _, o := range order.Others {
			if o.Orderer == c.Username {
				orders = append(orders, o)
			} else if o.Orderee == c.Username {
				orders = append(orders, o)
			}
		}
	}

	for _, order := range Charities {
		for _, o := range order.Orders {
			if o.Orderer == c.Username {
				orders = append(orders, o)
			} else if o.Orderee == c.Username {
				orders = append(orders, o)
			}
		}
	}

	for _, order := range Buisnesses {
		for _, o := range order.Orders {
			if o.Orderer == c.Username {
				orders = append(orders, o)
			} else if o.Orderee == c.Username {
				orders = append(orders, o)
			}
		}
	}

	return orders
}

func (r *Restaurant) GetOrders() []Order {
	var orders []Order = r.Others
	for _, order := range Charities {
		for _, o := range order.Orders {
			if o.Orderer == r.Username {
				orders = append(orders, o)
			} else if o.Orderee == r.Username {
				orders = append(orders, o)
			}
		}
	}

	for _, order := range Restaurants {
		for _, o := range order.Others {
			if o.Orderer == r.Username {
				orders = append(orders, o)
			} else if o.Orderee == r.Username {
				orders = append(orders, o)
			}
		}
	}

	for _, order := range Buisnesses {
		for _, o := range order.Orders {
			if o.Orderer == r.Username {
				orders = append(orders, o)
			} else if o.Orderee == r.Username {
				orders = append(orders, o)
			}
		}
	}

	return orders
}

func (b *Buisness) GetOrders() []Order {
	var orders []Order = b.Orders
	for _, order := range Charities {
		for _, o := range order.Orders {
			if o.Orderer == b.Username {
				orders = append(orders, o)
			}
		}
	}

	for _, order := range Restaurants {
		for _, o := range order.Others {
			if o.Orderer == b.Username {
				orders = append(orders, o)
			}
		}
	}

	return orders
}

func (r *Restaurant) GetReviews() []Review {
	return r.Reviews
}

func (u *User) GetUsers() []User {
	return Users.Users
}

func (c *Charity) Get(username string) *Charity {
	for _, charity := range Charities {
		if charity.Username == username {
			return &charity
		}
	}
	return nil
}

func (r *Restaurant) Get(username string) *Restaurant {
	for _, restaurant := range Restaurants {
		if restaurant.Username == username {
			return &restaurant
		}
	}
	return nil
}

func (u *User) Get(username string) *User {
	for _, user := range Users.Users {
		if user.Username == username {
			return &user
		}
	}
	return nil
}

func checkSessionToken(token string) (string, string, error) {
	collection := DB.Database("vwaste").Collection("users")
	for _, user := range Users.Users {
		if user.SessionToken == token {
			return user.Username, user.UserType, nil
		}
	}
	var user User
	err := collection.FindOne(context.Background(), bson.M{"session_token": token}).Decode(&user)
	return user.Username, user.SessionToken, err
}

func Authenticate(username, password string) (bool, string, error) {
	for _, user := range Charities {
		if user.Username == username && user.Password == password {
			return true, "charity", nil
		}
	}

	for _, user := range Restaurants {
		if user.Username == username && user.Password == password {
			return true, "restaurant", nil
		}
	}

	for _, user := range Buisnesses {
		if user.Username == username && user.Password == password {
			return true, "buisness", nil
		}
	}

	return false, "", nil
}

func (c *Charity) GetOrdersByType(orderType string) []Order {
	var orders []Order
	for _, order := range c.Orders {
		if order.OrderType == orderType {
			orders = append(orders, order)
		}
	}
	return orders
}

func (r *Restaurant) GetOrdersByType(orderType string) []Order {
	var orders []Order
	for _, order := range r.Others {
		if order.OrderType == orderType {
			orders = append(orders, order)
		}
	}
	return orders
}

func (c *Charity) GetOrdersByStatus(approved bool) []Order {
	var orders []Order
	for _, order := range c.Orders {
		if order.Approved == approved {
			orders = append(orders, order)
		}
	}
	return orders
}

func (r *Restaurant) GetOrdersByStatus(approved bool) []Order {
	var orders []Order
	for _, order := range r.Others {
		if order.Approved == approved {
			orders = append(orders, order)
		}
	}
	return orders
}

func bindSessionToken(token, username string) error {
	collection := DB.Database("vwaste").Collection("users")
	_, err := collection.UpdateOne(context.Background(), bson.M{"username": username}, bson.M{"$set": bson.M{"session_token": token}}, opts)

	for i, user := range Users.Users {
		if user.Username == username {
			Users.Users[i].SessionToken = token
			break
		}
	}
	return err
}

func __load_globalOrders() {
	var orders []Order
	for _, charity := range Charities {
		orders = append(orders, charity.Orders...)
	}

	for _, restaurant := range Restaurants {
		orders = append(orders, restaurant.Others...)
	}

	for _, buisness := range Buisnesses {
		orders = append(orders, buisness.Orders...)
	}

	// remove duplicates
	seen := make(map[int]bool)
	j := 0
	for _, order := range orders {
		if _, ok := seen[order.ID]; ok {
			continue
		}
		seen[order.ID] = true
		orders[j] = order
		j++
	}

	GlobalOrders = orders[:j]
}

func (c *Restaurant) AddEntry(name string, quantity int, _type string) error {
	fmt.Println("Adding Entry")
	collection := DB.Database("vwaste").Collection("charities")
	var waste = Waste{
		Product:  name,
		Quantity: quantity,
		Family:   _type,
	}

	_, err := collection.UpdateOne(context.Background(), bson.M{"username": c.Username}, bson.M{"$push": bson.M{"waste": waste}}, opts)

	if err != nil {
		var charity Restaurant
		collection.FindOne(context.Background(), bson.M{"username": c.Username}).Decode(&charity)
		charity.Waste = append(charity.Waste, waste)

		for i, charity := range Restaurants {
			if charity.Username == c.Username {
				Restaurants[i].Waste = append(Restaurants[i].Waste, waste)
				break
			}
		}

		collection.UpdateOne(context.Background(), bson.M{"username": c.Username}, bson.M{"$set": charity}, opts)
		return err
	}

	return nil
}

func (c *Buisness) AddEntry(name string, quantity int, _type string) error {
	collection := DB.Database("vwaste").Collection("buisnesses")
	var waste = Waste{
		Product:  name,
		Quantity: quantity,
		Family:   _type,
	}

	_, err := collection.UpdateOne(context.Background(), bson.M{"username": c.Username}, bson.M{"$push": bson.M{"waste": waste}}, opts)

	if err != nil {
		var buisness Buisness
		collection.FindOne(context.Background(), bson.M{"username": c.Username}).Decode(&buisness)
		buisness.Waste = append(buisness.Waste, waste)

		for i, buisness := range Buisnesses {
			if buisness.Username == c.Username {
				Buisnesses[i].Waste = append(Buisnesses[i].Waste, waste)
				break
			}
		}

		collection.UpdateOne(context.Background(), bson.M{"username": c.Username}, bson.M{"$set": buisness}, opts)
		return err
	}
	return nil
}
