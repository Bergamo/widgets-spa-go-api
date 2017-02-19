# go-rest-api
Simple REST API app with Go and MongoDB

# Dependencies
-go 1.8

# Instructions
- Clone the repository;
- Install project dependencies (To install all dependencies of a Golang project or golang projects recursively with the go getcommand, change directory into the project and simply run:):
```sh
$ go get ./...
```
- Build and run the application (change directory into the project)
```sh
$ go build
$ ./widgets-spa-go-api
```
- The app will start at port 3000

# GET TOKEN
- Makes a basic GET request to get a token

$ curl -http://localhost:3000/get-token

# GET USERS

$ curl http://localhost:3000/users -H 'Authorization: Bearer API_TOKEN'

# GET USER BY ID

$ curl http://localhost:3000/users/{id} -H 'Authorization: Bearer API_TOKEN'

# GET WIDGETS

$ curl http://localhost:3000/widgets -H 'Authorization: Bearer API_TOKEN'

# GET WIDGET BY ID

$ curl http://localhost:3000/widgets/{id} -H 'Authorization: Bearer API_TOKEN'

# CREATE WIDGET

$ curl -i -X POST "Content-Type:application/json" http://localhost:3000/widgets -d '{"name":"teste post","color":"magenta","price":"3.80","inventory":23,"melts":true}' -H 'Authorization: Bearer API_TOKEN'

# UPDATE WIDGET

$ curl -i -X PUT "Content-Type:application/json" http://localhost:3000/widgets/{id} -d '{"name":"teste","color":"magenta","price":"3.80","inventory":23,"melts":true}' -H 'Authorization: Bearer API_TOKEN'

#DELETE WIDGET

curl -i -X DELETE "Content-Type:application/json" http://localhost:3000/widgets/{id} -H 'Authorization: Bearer API_TOKEN'