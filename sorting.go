package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func SearchForWasteSources(location string, name string, filter string) []struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Username    string `json:"username"`
	Quantity    int    `json:"quantity"`
	Point       Point  `json:"point"`
} {

	var srcs []struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		Username    string `json:"username"`
		Quantity    int    `json:"quantity"`
		Point       Point  `json:"point"`
	}

	for _, restaurant := range Restaurants {
		isNot := false

		if location != "" && restaurant.Location != location {
			isNot = true
		}

		if name != "" && restaurant.Name != name {
			isNot = true
		}

		if filter != "" && restaurant.Description != filter {
			isNot = true
		}

		if isNot {
			continue
		}

		srcs = append(srcs, struct {
			Name        string `json:"name"`
			Description string `json:"description"`
			Username    string `json:"username"`
			Quantity    int    `json:"quantity"`
			Point       Point  `json:"point"`
		}{
			Name:        restaurant.Name,
			Description: restaurant.Description,
			Username:    restaurant.Username,
			Quantity:    len(restaurant.Waste),
			Point:       restaurant.Point,
		})

		if restaurant.Point == (Point{}) {
			srcs[len(srcs)-1].Point = Point{
				Rating:              3.5,
				LeaderboardPosition: 0,
			}
		}
	}

	for _, charity := range Charities {
		isNot := false

		if location != "" && charity.Location != location {
			isNot = true
		}

		if name != "" && charity.Name != name {
			isNot = true
		}

		if filter != "" && charity.Description != filter {
			isNot = true
		}

		if isNot {
			continue
		}
		srcs = append(srcs, struct {
			Name        string `json:"name"`
			Description string `json:"description"`
			Username    string `json:"username"`
			Quantity    int    `json:"quantity"`
			Point       Point  `json:"point"`
		}{
			Name:        charity.Name,
			Description: charity.Description,
			Username:    charity.Username,
			Quantity:    len(charity.Orders),
			Point:       Point{},
		})

	}

	for _, buisness := range Buisnesses {
		isNot := false

		if location != "" && buisness.Location != location {
			isNot = true
		}

		if name != "" && buisness.Name != name {
			isNot = true
		}

		if filter != "" && buisness.Description != filter {
			isNot = true
		}

		if isNot {
			continue
		}
		srcs = append(srcs, struct {
			Name        string `json:"name"`
			Description string `json:"description"`
			Username    string `json:"username"`
			Quantity    int    `json:"quantity"`
			Point       Point  `json:"point"`
		}{
			Name:        buisness.Name,
			Description: buisness.Description,
			Username:    buisness.Username,
			Quantity:    len(buisness.Waste),
			Point:       Point{},
		})

	}

	return srcs
}

func getGeoDecode(latlong string) (string, string, string) {
	url := "https://api.opencagedata.com/geocode/v1/json?q=%s+%s&key=03c48dae07364cabb7f121d8c1519492&no_annotations=1&language=en"
	lat := strings.Split(latlong, ",")[0]
	long := strings.Split(latlong, ",")[1]

	resp, err := http.Get(fmt.Sprintf(url, lat, long))
	if err != nil {
		return "", "", ""
	}

	defer resp.Body.Close()
	var res struct {
		Results []struct {
			Components struct {
				City     string `json:"city"`
				Road     string `json:"road"`
				Postcode string `json:"postcode"`
				Country  string `json:"country"`
			} `json:"components"`
		} `json:"results"`
	}

	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		return "", "", ""
	}

	return lat, long, fmt.Sprintf("%s, %s, %s, %s", res.Results[0].Components.Road, res.Results[0].Components.City, res.Results[0].Components.Postcode, res.Results[0].Components.Country)
}

func RequestWaste(order Order) error {
	OrderEEType := ""
	for _, restaurant := range Restaurants {
		if restaurant.Username == order.Orderee {
			OrderEEType = "restaurant"
			restaurant.AddOrder(order)
			break
		}
	}

	for _, charity := range Charities {
		if charity.Username == order.Orderee {
			OrderEEType = "charity"
			charity.AddOrder(order)
			break
		}
	}

	for _, buisness := range Buisnesses {
		if buisness.Username == order.Orderee {
			OrderEEType = "buisness"
			buisness.AddOrder(order)
			break
		}
	}

	if OrderEEType == "" {
		return fmt.Errorf("orderee not found")
	}

	if OrderEEType == "restaurant" {
		for _, restaurant := range Restaurants {
			if restaurant.Username == order.Orderee {
				order.Approved = false
				restaurant.Others = append(restaurant.Others, order)
			}
		}

		return nil
	}

	if OrderEEType == "charity" {
		for _, charity := range Charities {
			if charity.Username == order.Orderee {
				order.Approved = false
				charity.Orders = append(charity.Orders, order)
			}
		}

		return nil
	}

	if OrderEEType == "buisness" {
		for _, buisness := range Buisnesses {
			if buisness.Username == order.Orderee {
				order.Approved = false
				buisness.Orders = append(buisness.Orders, order)
			}
		}

		return nil
	}

	return fmt.Errorf("orderee not found")
}
