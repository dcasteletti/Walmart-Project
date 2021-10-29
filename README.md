# Walmart-Project

## Getting started

As you can see, this project was developed with `golang`; hence, before getting started, you need to follow the steps below.

1. Install [Golang](https://golang.org/).
2. Clone the [repository](https://github.com/dcasteletti/Walmart-Project.git) (and enter to the repository).
3. With the following commands you can download all the dependencies related to the project `go get -u all`. Or ...
4. This project uses MongoDB as database, therefore is necessary to run the service.
   1. Install [Brew](https://brew.sh/). 
   2. Install `brew services start mongodb-community@5.0`. 
   3. To run MongoDB if the connection was killed `brew services start mongodb-community@5.0`
5. [Load the products](https://github.com/walmartdigital/products-db.git) and follow the steps inside
6. Download [Robo 3T or Studio 3T](https://robomongo.org/). It is used to check the database
7. 


### MongoDb useful commands
+ `mongo` info 
+ `pkill mongod` kill the connection


### Lift the service
The API uses the 3001 port and boosts the main.go file. Uses the command below to lift the service
```
go run main.go
```

### Run unit tests

```
go test ./...
```

### In the file doc.yml, is the swagger of the API, it has the information of the developer for further information
