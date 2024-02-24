package main

import (
	"errors"
	"math/rand"
	"time"
)

func SignUpCharity(charity Charity) error {
	if charity.Username == "" || charity.Password == "" {
		return errors.New("all fields are required")
	}

	for _, c := range Charities {
		if c.Username == charity.Username {
			return errors.New("username already exists")
		}
	}

	Charities = append(Charities, charity)
	charity.Save()
	return nil
}

func SignUpRestaurant(restaurant Restaurant) error {
	if restaurant.Username == "" || restaurant.Password == "" {
		return errors.New("all fields are required")
	}

	for _, r := range Restaurants {
		if r.Username == restaurant.Username {
			return errors.New("username already exists")
		}
	}

	Restaurants = append(Restaurants, restaurant)
	restaurant.Save()
	return nil
}

func SignUpBuisness(buisness Buisness) error {
	if buisness.Username == "" || buisness.Password == "" {
		return errors.New("all fields are required")
	}

	for _, b := range Buisnesses {
		if b.Username == buisness.Username {
			return errors.New("username already exists")
		}
	}

	Buisnesses = append(Buisnesses, buisness)
	buisness.Save()
	return nil
}

func generateSessionToken() string {
	rand := rand.New(rand.NewSource(time.Now().UnixNano()))
	ascii := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	token := make([]byte, 32)
	for i := range token {
		token[i] = ascii[rand.Intn(len(ascii))]
	}

	return string(token)
}

func SignIn(username, password string) (bool, string, string, error) {
	auth, _type, err := Authenticate(username, password)
	if err != nil {
		return false, "", "", err
	}

	if !auth {
		return false, "", "", errors.New("invalid credentials")
	}

	user := &User{
		Username: username,
		Password: password,
		UserType: _type,
	}

	user.Save()
	for i, u := range Users.Users {
		if u.Username == username {
			Users.Users[i].SessionToken = generateSessionToken()
			return true, _type, Users.Users[i].SessionToken, nil
		}
	}

	sessionToken := generateSessionToken()
	bindSessionToken(sessionToken, username)

	return true, _type, sessionToken, nil
}

func GetTokenUsername(token string) (string, string, error) {
	username, userType, err := checkSessionToken(token)
	if err != nil {
		return "", "", errors.New("invalid token, please sign in again")
	}

	return username, userType, nil
}
