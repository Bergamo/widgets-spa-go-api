# widgets-spa-go-api
RESTful API with Golang and MongoDB and Authentication with  JSON Web Tokens.

# Dependencies
- go^1.8
- MongoDB^3.4.2

# Install the Go tools
- To install go tools visit : https://golang.org/doc/install

# Instructions
- Clone the repository inside the your go workspace (example: work/src/github.com/Bergamo/)
- Change directory into the project and Install project dependencies
```sh
$ go get ./...
```
- Build and run the application
```sh
$ go build
```
- Run the application (Make sure the mongo is running)

### MacOS and Linux
```sh
$ ./widgets-spa-go-api
```
### Windows
```sh
$ go run server.go
```

- The app will start at port 3000

# The endpoints are as follows:

- GET /get-token http://localhost:3000/get-token
- GET /users http://localhost:3000/users
- GET /users/:id http://localhost:3000/users/:id
- GET /widgets http://localhost:3000/widgets
- GET /widgets/:id http://localhost:3000/widgets/:id
- POST /widgets for creating new widgets http://localhost:3000/widgets
- PUT /widgets/:id for updating existing widgets http://localhost:3000/widgets/:id

# API Tests

## GET TOKEN
- Makes a basic GET request to get a token

```sh
$ curl http://localhost:3000/get-token
```

## GET USERS
- Change API_TOKEN for the token obtained in GET TOKEN request
```sh
$ curl http://localhost:3000/users -H 'Authorization: Bearer API_TOKEN'
```

## GET USER BY ID
- Change API_TOKEN for the token obtained in GET TOKEN request
- Change :id for the Widget id
```sh
$ curl http://localhost:3000/users/:id -H 'Authorization: Bearer API_TOKEN'
```

## GET WIDGETS
- Change API_TOKEN for the token obtained in GET TOKEN request
```sh
$ curl http://localhost:3000/widgets -H 'Authorization: Bearer API_TOKEN'
```

## GET WIDGET BY ID
- Change API_TOKEN for the token obtained in GET TOKEN request
- Change :id for the Widget id
```sh
$ curl http://localhost:3000/widgets/:id -H 'Authorization: Bearer API_TOKEN'
```

## CREATE WIDGET
- Change API_TOKEN for the token obtained in GET TOKEN request
```sh
$ curl -i -X POST "Content-Type:application/json" http://localhost:3000/widgets -d '{"name":"teste post","color":"magenta","price":"3.80","inventory":23,"melts":true}' -H 'Authorization: Bearer API_TOKEN'
```

## UPDATE WIDGET
- Change API_TOKEN for the token obtained in GET TOKEN request
- Change :id for the Widget id
```sh
$ curl -i -X PUT "Content-Type:application/json" http://localhost:3000/widgets/:id -d '{"name":"teste","color":"magenta","price":"3.80","inventory":23,"melts":true}' -H 'Authorization: Bearer API_TOKEN'
```

## DELETE WIDGET
- Change API_TOKEN for the token obtained in GET TOKEN request
- Change :id for the Widget id
```sh
$ curl -i -X DELETE "Content-Type:application/json" http://localhost:3000/widgets/:id -H 'Authorization: Bearer API_TOKEN'
```