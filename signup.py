from requests import post


json_Data_signup = {
    "name": "Dominos",
    "location": "33,45",
    "username": "dom",
    "password": "dom",\
    "description": "Pizza and more, Full menu with prices, Fast & easy online ordering, Delivery, Pickup, Catering, and more, Open late.",
}

json_Data = {
    "id": 1299,
    "product": "Pizza",
    "quantity": 8,
    "orderer": "dom",
    "orderee": "paragon",
}

response = post("http://localhost:8080/api/signup/restaurant", json=json_Data_signup)
print(response.text)
