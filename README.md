# Majoo Golang Bootcamp

## Setup project
 - Clone project
 - Run command
    ```
    go mod tidy
    ```

## Run project

 - Go to root folder
 - Run command
    ```
    go run ./cmd/main.go
    ```
- Server will run in `localhost:8000`

## API List
- Auth
   - POST v1/login
   - POST v1/register
   
 - User
    - POST v1/user
    - PUT v1/user
    - DELETE v1/user/:ID
 - Outlet
    - GET v1/outlet
    - GET v1/outlet/:ID
    - POST v1/outlet
    - PUT v1/outlet
    - DELETE v1/outlet/:ID
 - Category
 - Product
    - GET v1/products
    - GET v1/products/:ID
    - POST v1/products
    - PUT v1/products
    - DELETE v1/products/:ID
 - Transaction
    - GET v1/transaction
    - GET v1/transaction/:ID
    - POST v1/transaction
    - PUT v1/transaction
    - POST v1/transaction/payment
    
## Document
- ERD
![logo](https://github.com/upgradeskill/fp2022-majoo-djaya-bersama/blob/main/mini-pos.png)
