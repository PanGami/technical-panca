# Technical Panca API
This application is built to fulfill the technical test requirements provided by Synapsis using the following tech stack: Golang, Fiber, Gorm, and PostgreSQL, running in Docker. It includes features ranging from user registration, login, product list by category, shopping cart, list cart, checkout, etc.

# Prerequisite
`` git ``
`` docker ``

# Instalation 

- Clone repo : ```https://github.com/PanGami/technical-panca.git```
- Go inside project: ```'cd technical-panca'```
- Run Docker : ```docker-compose up --build```

# Endpoint REST API Index

## Conn Check
- **Check Redis**
  - **Method:** GET
  - **URL:** localhost:8000/api/check/redis

- **Check Posgres**
  - **Method:** GET
  - **URL:** localhost:8000/api/check/postgres

## Auth
- **Register**
  - **Method:** POST
  - **URL:** localhost:8000/api/auth/register

- **Login**
  - **Method:** POST
  - **URL:** localhost:8000/api/auth/login

- **Logout**
  - **Method:** GET
  - **URL:** localhost:8000/api/auth/logout

- **Refresh**
  - **Method:** GET
  - **URL:** localhost:8000/api/auth/refresh

## User
- **Logged User**
  - **Method:** GET
  - **URL:** localhost:8000/api/users/me

## Product
- **Get all products**
  - **Method:** GET
  - **URL:** localhost:8000/api/products

- **Get Product**
  - **Method:** GET
  - **URL:** localhost:8000/api/products/{{ProductID}}

- **Add Products**
  - **Method:** POST
  - **URL:** localhost:8000/api/products

- **Update Product**
  - **Method:** PUT
  - **URL:** localhost:8000/api/products/{{ProductID}}

- **Delete Product**
  - **Method:** DELETE
  - **URL:** localhost:8000/api/products/{{ProductID}}

- **Get By category**
  - **Method:** GET
  - **URL:** localhost:8000/api/products/category?category={{Category}}

## Cart
- **Add Cart Item**
  - **Method:** POST
  - **URL:** localhost:8000/api/cart/add

- **Get Cart Items**
  - **Method:** GET
  - **URL:** localhost:8000/api/cart/items?user_id={{UserId}}

- **Delete Cart Item**
  - **Method:** DELETE
  - **URL:** localhost:8000/api/cart/delete/{{CartID}}

- **Checkout**
  - **Method:** POST
  - **URL:** localhost:8000/api/cart/checkout

## Test Default (If not found)
- **Method:** GET
- **URL:** localhost:8000/